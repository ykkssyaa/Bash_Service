package gateway

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type Gateways struct {
	Command Command
	Storage Storage
}

func NewGateway(db *sqlx.DB) *Gateways {
	return &Gateways{
		Command: NewCommandPostgres(db),
		Storage: NewCtxStorage(),
	}
}

type Storage interface {
	Get(id int) context.CancelFunc
	Set(id int, ctxFunc context.CancelFunc)
	Remove(id int)
}
