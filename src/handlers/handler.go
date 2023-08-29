package handlers

import (
	"context"
	pb "go_unit_test/pkg/pb"
	"go_unit_test/src/entities"
	"go_unit_test/src/usecases/book"
)

type BookHandler struct {
	getUsecase book.GetUsecase
}

func New(getUsecase book.GetUsecase) pb.BookServiceServer {
	return &BookHandler{getUsecase: getUsecase}
}
func (h *BookHandler) Get(ctx context.Context, in *pb.GetRequest) (*pb.Book, error) {
	filter := entities.BookFillter{
		Title: in.Title,
	}
	book, err := h.getUsecase.Execute(ctx, filter)
	if err != nil {
		return nil, err
	}

	response := &pb.Book{
		Title:  book.Title,
		Author: book.Author,
	}
	return response, nil
}
