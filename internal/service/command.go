package service

import "github.com/ykkssyaa/Bash_Service/internal/gateway"
import lg "github.com/ykkssyaa/Bash_Service/pkg/logger"

type Command interface {
}

type CommandService struct {
	repo   gateway.Command
	logger *lg.Logger
}

func NewCommandService(repo gateway.Command, logger *lg.Logger) *CommandService {
	return &CommandService{repo: repo, logger: logger}
}
