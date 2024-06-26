package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/ykkssyaa/Bash_Service/internal/consts"
	"github.com/ykkssyaa/Bash_Service/internal/gateway"
	"github.com/ykkssyaa/Bash_Service/internal/models"
	se "github.com/ykkssyaa/Bash_Service/pkg/serverError"
	"net/http"
)
import lg "github.com/ykkssyaa/Bash_Service/pkg/logger"

type Executor interface {
	ExecCmd(ctx context.Context, script string, ch <-chan int) error
}

type CommandService struct {
	repo       gateway.Command
	ctxStorage gateway.Storage
	cache      gateway.Cache
	executor   Executor
	logger     *lg.Logger
}

func NewCommandService(repo gateway.Command, ctxStorage gateway.Storage,
	cache gateway.Cache, executor Executor, logger *lg.Logger) *CommandService {
	return &CommandService{repo: repo, logger: logger, ctxStorage: ctxStorage, cache: cache, executor: executor}
}

func (c CommandService) CreateCommand(script string) (models.Command, error) {

	if script == "" {
		return models.Command{}, se.ServerError{
			Message:    consts.ErrorEmptyScript,
			StatusCode: http.StatusBadRequest,
		}
	}

	cmd := models.Command{Script: script, Status: models.StatusStarted}
	ch := make(chan int, 1) // Канал для передачи id сохраненной команды

	ctx, cancel := context.WithTimeout(context.Background(), consts.CtxTimeout)

	err := c.executor.ExecCmd(ctx, script, ch)
	if err != nil {
		c.logger.Err.Println(consts.ErrorCreateCommand, err.Error())

		cancel() // Cancel ctx

		return models.Command{}, se.ServerError{
			Message:    consts.ErrorExecCommand,
			StatusCode: http.StatusInternalServerError,
		}
	}

	id, err := c.repo.CreateCommand(cmd)
	if err != nil {
		c.logger.Err.Println(consts.ErrorCreateCommand, err.Error())

		cancel() // Cancel ctx

		return models.Command{}, se.ServerError{
			Message:    consts.ErrorCreateCommand,
			StatusCode: http.StatusInternalServerError,
		}
	}

	// Save ctx cancel func in storage
	c.ctxStorage.Set(id, cancel)

	cmd.Id = id
	ch <- id

	return cmd, nil
}

func (c CommandService) GetCommand(id int) (models.Command, error) {

	if id <= 0 {
		return models.Command{}, se.ServerError{
			Message:    consts.ErrorWrongId,
			StatusCode: http.StatusBadRequest,
		}
	}

	// Ищем команду в кеше
	command, err := c.cache.Get(id)
	if err != nil {
		c.logger.Err.Println(consts.ErrorGetCommand, err.Error())
	}
	// Если команда не была в кеше, обращаемся к БД
	if command.Id == 0 {
		command, err = c.repo.GetCommand(id)
		if err != nil {
			c.logger.Err.Println(consts.ErrorGetCommand, err.Error())

			if errors.Is(err, sql.ErrNoRows) {
				return models.Command{}, se.ServerError{
					Message:    "",
					StatusCode: http.StatusNotFound,
				}
			}

			return models.Command{}, se.ServerError{
				Message:    consts.ErrorGetCommand,
				StatusCode: http.StatusInternalServerError,
			}
		}
	}

	return command, err
}

func (c CommandService) GetAllCommands(limit, offset int) ([]models.Command, error) {

	if limit < 0 {
		return nil, se.ServerError{
			Message:    consts.ErrorWrongLimit,
			StatusCode: http.StatusBadRequest,
		}
	}

	if offset < 0 {
		return nil, se.ServerError{
			Message:    consts.ErrorWrongOffset,
			StatusCode: http.StatusBadRequest,
		}
	}

	commands, err := c.repo.GetAllCommands(limit, offset)
	if err != nil {
		c.logger.Err.Println(consts.ErrorGetCommand, err.Error())
		return nil, se.ServerError{
			Message:    consts.ErrorGetCommand,
			StatusCode: http.StatusInternalServerError,
		}
	}

	return commands, err
}

func (c CommandService) StopCommand(id int) error {

	if id <= 0 {
		return se.ServerError{
			Message:    consts.ErrorWrongId,
			StatusCode: http.StatusBadRequest,
		}
	}

	cancel := c.ctxStorage.Get(id)

	if cancel == nil {
		c.logger.Err.Printf("cancel to command %d not found\n", id)
		return se.ServerError{
			Message:    "",
			StatusCode: http.StatusNotFound,
		}
	}

	cancel()

	return nil
}
