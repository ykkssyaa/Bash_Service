package service

import (
	"bufio"
	"context"
	"errors"
	"github.com/ykkssyaa/Bash_Service/internal/consts"
	"github.com/ykkssyaa/Bash_Service/internal/models"
	"io"
	"os/exec"
	"time"
)

// ExecCmd создает процесс bash-скрипта CommandContext, который задается аргументом script
// Контекст передается для возможности отмены выполнения скрипта.
// В конце работы процесса из хранилища удаляется функция отмены контекста
// В ch передается id созданной записи в БД о выполняемой команде
func (c CommandService) ExecCmd(ctx context.Context, script string, ch <-chan int) error {

	status := models.StatusStarted

	// Инициализация процесса скрипта
	cmd := exec.CommandContext(ctx, "bash", "-c", script)

	// Захват вывода скрипта
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	// Старт процесса
	if err := cmd.Start(); err != nil {
		return err
	}

	// Обработка вывода скрипта в отдельной горутине
	go func() {

		// Буфер записи вывода скрипта
		var outputBuffer []byte
		// Канал для передачи статуса работы процесса
		statusCh := make(chan string)
		// Reader для вывода процесса
		stdoutReader := bufio.NewReader(stdout)

		// Запуск чтения из stdout процесса
		go readOutput(stdoutReader, &outputBuffer, statusCh)

		// Таймер для обновления состояния программы
		ticker := time.NewTicker(consts.ReadOutputTime)
		defer ticker.Stop()

		// Ожидание создания записи о команде (передается id записи)
		var id int
		id = <-ch

		// Объект команды для дальнейшего сохранения
		cm := models.Command{
			Id:     id,
			Output: "",
			Status: status,
			Script: script,
		}

	Loop:
		for { // Цикл, пока не закончится вывод процесса
			select {
			case <-ticker.C: // При каждом тике сохраняем вывод из буфера и обновляем данные о процессе
				cm.Output = string(outputBuffer)

				// Кешируем данное значение для уменьшения количества обращений к БД
				if err := c.cache.Set(id, cm); err != nil {
					c.logger.Err.Println(consts.ErrorUpdateCacheCommand, err.Error())
				}

				//if err = c.repo.UpdateCommand(cm); err != nil {
				//	c.logger.Err.Println(consts.ErrorUpdateCommand, err.Error())
				//}

			case st := <-statusCh: // При получении статуса завершаем цикл и сохраняем статус
				cm.Status = st
				break Loop
			}
		}

		// Ожидание завершения процесса
		if err := cmd.Wait(); err != nil {
			if errors.Is(ctx.Err(), context.Canceled) { // Если был отменён контекст
				status = models.StatusStopped
			} else { // При какой-либо ошибке
				c.logger.Err.Println("process error: ", err)
				status = models.StatusError
			}

		} else {
			status = models.StatusSuccess
		}

		cm.Status = status

		// Обновление данных
		if err = c.repo.UpdateCommand(cm); err != nil {
			c.logger.Err.Println(consts.ErrorUpdateCommand, err.Error())
		}

		// Удаляем из кеша команду
		if err := c.cache.Remove(id); err != nil {
			c.logger.Err.Println(consts.ErrorRemoveCacheCommand, err.Error())
		}

		// Удаление записи о контексте
		c.ctxStorage.Remove(id)
	}()

	return nil
}

// readOutput читает построчно из stdout и записывает в буфер outputBuffer.
// При завершении отправляет в канал статус завершения процесса: ошибка/успех
func readOutput(stdout *bufio.Reader, outputBuffer *[]byte, statusCh chan<- string) {
	for {
		line, _, err := stdout.ReadLine()
		if err != nil {
			if err == io.EOF {
				statusCh <- models.StatusSuccess
			} else {
				statusCh <- models.StatusError
			}
			break
		}
		*outputBuffer = append(*outputBuffer, line...)
		*outputBuffer = append(*outputBuffer, '\n')
	}
}
