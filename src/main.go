package main

import (
	"fmt"
	pb "go_unit_test/pkg/pb"
	"go_unit_test/src/gateways/repository/book"
	"go_unit_test/src/handlers"
	bookusecase "go_unit_test/src/usecases/book"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	bookRepository := book.NewRepository()
	getUsecase := bookusecase.New(bookRepository)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterBookServiceServer(s, handlers.New(getUsecase))
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	fmt.Println(getUsecase)
}
