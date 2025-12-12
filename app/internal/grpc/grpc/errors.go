package grpc

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r AuthServer) ErrInvalidArgument(err error) error {
	message := "invalid argument"

	if err != nil {
		message = err.Error()
	}

	return status.Error(codes.InvalidArgument, message)
}

func (r AuthServer) ErrInternal(err error) error {
	message := "internal error"

	if err != nil {
		message = err.Error()
	}

	return status.Error(codes.Internal, message)
}
