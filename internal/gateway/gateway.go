package gateway

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/ykkssyaa/Bash_Service/internal/models"
)

type Gateways struct {
	Command      Command
	Storage      Storage
	CommandCache Cache
}

func NewGateway(db *sqlx.DB, storageSize int) *Gateways {
	return &Gateways{
		Command:      NewCommandPostgres(db),
		Storage:      NewCtxStorage(storageSize),
		CommandCache: NewCommandCache(storageSize),
	}
}

type Storage interface {
	Get(id int) context.CancelFunc
	Set(id int, ctxFunc context.CancelFunc)
	Remove(id int)
}

type Cache interface {
	Get(id int) (models.Command, error)
	Set(id int, cmd models.Command) error
	Remove(id int) error
}

type Command interface {
	CreateCommand(command models.Command) (id int, err error)
	UpdateCommand(command models.Command) error
	GetCommand(commandId int) (models.Command, error)
	GetAllCommands(limit, offset int) ([]models.Command, error)
}
