package service

import (
	"context"
	"github.com/ykkssyaa/Bash_Service/internal/consts"
	"github.com/ykkssyaa/Bash_Service/internal/gateway"
	"github.com/ykkssyaa/Bash_Service/internal/models"
	se "github.com/ykkssyaa/Bash_Service/pkg/serverError"
	"net/http"
)
import lg "github.com/ykkssyaa/Bash_Service/pkg/logger"

type Command interface {
	CreateCommand(script string) (models.Command, error)
	GetCommand(id int) (models.Command, error)
	GetAllCommands(limit, offset int) ([]models.Command, error)
}

type CommandService struct {
	repo       gateway.Command
	ctxStorage gateway.Storage
	logger     *lg.Logger
}

func NewCommandService(repo gateway.Command, logger *lg.Logger) *CommandService {
	return &CommandService{repo: repo, logger: logger}
}

func (c CommandService) CreateCommand(script string) (models.Command, error) {

	cmd := models.Command{Script: script, Status: models.StatusStarted}
	ch := make(chan int, 1) // Канал для передачи id сохраненной команды

	ctx, cancel := context.WithTimeout(context.Background(), consts.CtxTimeout)

	go c.ExecCmd(ctx, script, ch)

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
	go c.ctxStorage.Set(id, cancel)

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

	return c.repo.GetCommand(id)
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

	return c.repo.GetAllCommands(limit, offset)
}
