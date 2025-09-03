package grpc

import (
	atomWebsite "github.com/dzamyatin/atomWebsite/internal/grpc/generated"
)

type ShopServer struct {
	atomWebsite.UnimplementedShopServer
}

func NewShopServer() ShopServer {
	return ShopServer{}
}
