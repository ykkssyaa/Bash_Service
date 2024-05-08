package gateway

import (
	"github.com/jmoiron/sqlx"
	"github.com/ykkssyaa/Bash_Service/internal/models"
)

type Command interface {
	CreateCommand(command models.Command) (id int, err error)
	UpdateCommand(command models.Command) error
}

type CommandPostgres struct {
	db *sqlx.DB
}

func NewCommandPostgres(db *sqlx.DB) *CommandPostgres {
	return &CommandPostgres{db: db}
}

func (c CommandPostgres) CreateCommand(command models.Command) (id int, err error) {

	query := "INSERT INTO commands (script, status, output) VALUES ($1, $2, $3) RETURNING id"

	tx, err := c.db.Begin()

	if err != nil {
		return 0, err
	}

	row := tx.QueryRow(query, command.Script, command.Status, command.Output)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (c CommandPostgres) UpdateCommand(command models.Command) error {

	query := "UPDATE commands SET status=$1, output=$2 WHERE id=$3"

	tx, err := c.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(query, command.Status, command.Output, command.Id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
