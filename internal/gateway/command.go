package gateway

import "github.com/jmoiron/sqlx"

type Command interface {
}

type CommandPostgres struct {
	db *sqlx.DB
}

func NewCommandPostgres(db *sqlx.DB) *CommandPostgres {
	return &CommandPostgres{db: db}
}
