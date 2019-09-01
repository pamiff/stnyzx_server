package service

import (
	"github.com/pamiff/stnyzx_server/auth/proto"
	//"github.com/micro/go-micro/config"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/service/kubernetes"
	"github.com/pamiff/stnyzx_server/auth/config"
	"github.com/pamiff/stnyzx_server/auth/dao"
)

type Service struct {
	service micro.Service
	config  *config.Config
}

func New(d *dao.Dao, c *config.Config) *Service {
	var s micro.Service
	if config.ProgramEnvironment == "dev" {
		s = micro.NewService(
			micro.Name(config.ServiceName),
		)
	} else if config.ProgramEnvironment == "test" ||
		config.ProgramEnvironment == "prod" {
		s = kubernetes.NewService(
			micro.Name(config.ServiceName),
		)
	}
	s.Init()
	handler := NewAuthHandlerImpl(d, s.Client(), c)
	proto.RegisterAuthHandler(s.Server(), handler)
	return &Service{
		service: s,
		config:  c,
	}
}

func (s *Service) Run() error {
	return s.service.Run()
}
