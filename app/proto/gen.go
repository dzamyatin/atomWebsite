package proto

//go:generate protoc -I. -I./google --openapiv2_out . --go_out=../internal/grpc/generated/ --go_opt=paths=source_relative --go-grpc_out=../internal/grpc/generated/ --go-grpc_opt=paths=source_relative auth.proto
