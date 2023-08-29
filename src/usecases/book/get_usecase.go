package book

import (
	"context"
	"go_unit_test/src/entities"
	"go_unit_test/src/gateways/repository/book"
)

type getUsecase struct {
	repository book.BookRepository
}

func New(repository book.BookRepository) GetUsecase {
	return getUsecase{repository: repository}
}

func (g getUsecase) Execute(ctx context.Context, filter entities.BookFillter) (entities.Book, error) {
	return g.repository.Get(ctx, filter)
}
