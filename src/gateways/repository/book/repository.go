package book

import (
	"context"
	"errors"
	"go_unit_test/src/entities"
)

type bookRepository struct{}

var (
	books = []entities.Book{
		{Title: "The colors", Author: "J. April"},
		{Title: "The flowers", Author: "O. Kyle"},
		{Title: "The cars", Author: "P. Joseph"},
	}
)

func NewRepository() BookRepository {
	return bookRepository{}
}

func (b bookRepository) Get(ctx context.Context, filter entities.BookFillter) (entities.Book, error) {
	for _, book := range books {
		if book.Title == filter.Title {
			return book, nil
		}
	}
	return entities.Book{}, errors.New("not found")
}
