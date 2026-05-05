package repository

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const driverName = "pgx"

type TaskContext struct {
	Dbx *sqlx.DB
}

func NewTaskContext() *TaskContext {
	return &TaskContext{}
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "taskmanagerdb"
)

func (s *TaskContext) Connect() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	dbx, err := sqlx.Open("postgres", psqlInfo)

	if err != nil {
		errors.New("erro ao conectar no DB!")
		return err
	}
	s.Dbx = dbx
	return nil
}

func (s *TaskContext) Close() error {
	return s.Dbx.Close()
}
