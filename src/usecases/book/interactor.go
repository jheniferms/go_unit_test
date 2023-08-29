package book

import (
	"context"
	"go_unit_test/src/entities"
)

type GetUsecase interface {
	Execute(ctx context.Context, filter entities.BookFillter) (entities.Book, error)
}
