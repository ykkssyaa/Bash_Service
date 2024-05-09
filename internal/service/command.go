package service

import (
	"github.com/ykkssyaa/Bash_Service/internal/consts"
	"github.com/ykkssyaa/Bash_Service/internal/gateway"
	"github.com/ykkssyaa/Bash_Service/internal/models"
	se "github.com/ykkssyaa/Bash_Service/pkg/serverError"
	"net/http"
)
import lg "github.com/ykkssyaa/Bash_Service/pkg/logger"

type Command interface {
	CreateCommand(script string) (models.Command, error)
}

type CommandService struct {
	repo   gateway.Command
	logger *lg.Logger
}

func NewCommandService(repo gateway.Command, logger *lg.Logger) *CommandService {
	return &CommandService{repo: repo, logger: logger}
}

func (c CommandService) CreateCommand(script string) (models.Command, error) {

	cmd := models.Command{Script: script, Status: models.StatusStarted}
	ch := make(chan int, 1) // Канал для передачи id сохраненной команды

	go c.ExecCmd(script, ch)

	id, err := c.repo.CreateCommand(cmd)
	if err != nil {
		c.logger.Err.Println(consts.ErrorCreateCommand, err.Error())
		return models.Command{}, se.ServerError{
			Message:    consts.ErrorCreateCommand,
			StatusCode: http.StatusInternalServerError,
		}
	}
	cmd.Id = id
	ch <- id

	return cmd, nil
}
