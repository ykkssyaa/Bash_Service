package gateway

import "github.com/jmoiron/sqlx"

type Gateways struct {
	Command Command
}

func NewGateway(db *sqlx.DB) *Gateways {
	return &Gateways{
		Command: NewCommandPostgres(db),
	}
}
