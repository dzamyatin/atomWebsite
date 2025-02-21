package proto

//go:generate sh -c "find ./ -name '*.proto' | xargs protoc --go_out=../internal/grpc/generated/ --go_opt=paths=source_relative --go-grpc_out=../internal/grpc/generated/ --go-grpc_opt=paths=source_relative"
