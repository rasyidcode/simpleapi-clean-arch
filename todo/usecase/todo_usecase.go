package usecase

import (
	"context"
	"time"

	"rasyidcode/simpleapi-clean-arch/models"
)

type todoUsecase struct {
	todoRepo       models.TodoRepository
	contextTimeout time.Duration
}

func NewTodoUsecase(tr models.TodoRepository, timeout time.Duration) models.TodoUsecase {
	return &todoUsecase{
		todoRepo:       tr,
		contextTimeout: timeout,
	}
}

func (t *todoUsecase) GetByID(c context.Context, id int64) (res models.Todo, err error) {
	ctx, cancel := context.WithTimeout(c, t.contextTimeout)
	defer cancel()

	res, err = t.todoRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	return
}
