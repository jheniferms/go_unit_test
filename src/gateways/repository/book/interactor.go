package book

import (
	"context"
	"go_unit_test/src/entities"
)

type BookRepository interface {
	Get(ctx context.Context, filter entities.BookFillter) (entities.Book, error)
}
