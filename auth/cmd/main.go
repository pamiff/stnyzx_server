package main

import (
	"fmt"
	"github.com/pamiff/stnyzx_server/auth/config"
	"github.com/pamiff/stnyzx_server/auth/dao"
	"github.com/pamiff/stnyzx_server/auth/service"
	// "github.com/micro/go-micro"
	// "github.com/micro/go-micro/service/grpc"
	// // k8s "github.com/micro/kubernetes/go/micro"
	// //k8s "github.com/micro/kubernetes/go/micro"
	// proto "github.com/pamiff/stnyzx_server
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	d := dao.New(&cfg)
	s := service.New(d, &cfg)
	s.Run()
	// service := grpc.NewService(
	// 	micro.Name("auth"),
	// )
	// service.Init()
	// proto.RegisterAuthHandler(service.Server(), new(proto.AuthHandlerImpl))
	// service.Run()
}
