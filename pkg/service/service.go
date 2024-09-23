package service

import (
	"go.uber.org/zap"
	"todo"
	"todo/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
}

type TodoList interface{}

type TodoItem interface{}

type Service struct {
	Authorization
	TodoList
	TodoItem
	log zap.Logger
}

func NewService(repos *repository.Repository, log zap.Logger) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		log:           log,
	}
}
