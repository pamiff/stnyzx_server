package main

import (
	"github.com/pamiff/stnyzx_server/auth/proto"
	"context"
	"github.com/micro/go-micro"
	// "github.com/micro/go-micro/service/grpc"
	// "github.com/pamiff/stnyzx_server/auth/"
)

func main() {
	service := micro.NewService(
		micro.Name("apigateway"),
	)
	authClient := proto.NewAuthService("auth", service.Client())
	authClient.GetJwt(context.Background(), &proto.GetJwtRequest{Code: "test"})
}
