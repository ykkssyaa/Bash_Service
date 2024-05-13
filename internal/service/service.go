package service

import (
	"github.com/ykkssyaa/Bash_Service/internal/gateway"
	"github.com/ykkssyaa/Bash_Service/internal/models"
	"github.com/ykkssyaa/Bash_Service/pkg/logger"
)

type Services struct {
	Command Command
}

func NewService(gateways *gateway.Gateways, logger *logger.Logger) *Services {
	return &Services{
		Command: NewCommandService(gateways.Command, gateways.Storage, gateways.CommandCache, logger),
	}
}

type Command interface {
	CreateCommand(script string) (models.Command, error)
	GetCommand(id int) (models.Command, error)
	GetAllCommands(limit, offset int) ([]models.Command, error)
	StopCommand(id int) error
}
