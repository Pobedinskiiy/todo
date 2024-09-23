package repository

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"todo"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
}

type TodoList interface{}

type TodoItem interface{}

type Repository struct {
	Authorization
	TodoList
	TodoItem
	log zap.Logger
}

func NewRepository(db *sqlx.DB, log zap.Logger) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		log:           log,
	}
}
