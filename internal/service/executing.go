package service

import (
	"context"
	"github.com/ykkssyaa/Bash_Service/internal/consts"
	"github.com/ykkssyaa/Bash_Service/internal/models"
	"os/exec"
)

func (c CommandService) ExecCmd(ctx context.Context, script string, ch <-chan int) {

	cmd := exec.CommandContext(ctx, "bash", "-c", script)

	output, err := cmd.CombinedOutput()

	var id int
	id = <-ch

	var status string
	if err != nil { // При возникновении ошибки, меняем статус на ошибку
		status = models.StatusError
	} else {
		status = models.StatusSuccess
	}

	err = c.repo.UpdateCommand(
		models.Command{
			Id:     id,
			Output: string(output),
			Status: status,
			Script: script,
		})

	if err != nil {
		c.logger.Err.Println(consts.ErrorUpdateCommand, err.Error())
	}

	// Delete ctx cancel func
	go c.ctxStorage.Remove(id)
}
