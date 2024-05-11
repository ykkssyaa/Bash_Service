package service

import (
	"github.com/ykkssyaa/Bash_Service/internal/gateway"
	"github.com/ykkssyaa/Bash_Service/pkg/logger"
)

type Services struct {
	Command Command
}

func NewService(gateways *gateway.Gateways, logger *logger.Logger) *Services {
	return &Services{
		Command: NewCommandService(gateways.Command, gateways.Storage, logger),
	}
}
