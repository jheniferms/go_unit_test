run:
	go run src/main.go

build-proto:
	protoc ./pkg/book.proto --go_out=./pkg --go-grpc_out=./pkg --go-grpc_opt=require_unimplemented_servers=false
	