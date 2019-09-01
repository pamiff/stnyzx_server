package service

import (
	"github.com/pamiff/stnyzx_server/auth/config"
	"github.com/micro/go-micro/client"
	"github.com/pamiff/stnyzx_server/auth/dao"
	"github.com/pamiff/stnyzx_server/auth/proto"
	"context"
	"fmt"
)

type AuthHandlerImpl struct {
	proto.AuthHandler
	dao *dao.Dao
	cfg *config.Config
}

func NewAuthHandlerImpl(d *dao.Dao, cli client.Client, cfg *config.Config) *AuthHandlerImpl {
	return &AuthHandlerImpl {
		dao: d,
		cfg: cfg,
	}
}

func (h *AuthHandlerImpl) GetJwt(c context.Context, resquest *proto.GetJwtRequest, response *proto.GetJwtResponse) error {
	fmt.Println(resquest.Code)
	response.Jwt = "test"
	return nil
}
