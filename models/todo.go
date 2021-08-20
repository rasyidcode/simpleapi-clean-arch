package models

import (
	"context"
)

type Todo struct {
	ID       int64  `json:"id"`
	TodoName string `json:"todo_name"`
	Desc     string `json:"desc"`
}

type TodoUsecase interface {
	GetByID(ctx context.Context, id int64) (Todo, error)
}

type TodoRepository interface {
	GetByID(ctx context.Context, id int64) (Todo, error)
}
