package dao

import (
	"github.com/pamiff/stnyzx_server/auth/config"
)

type Dao struct {
	config *config.Config
}

func New(c *config.Config) *Dao {
	return &Dao{
		config: c,
	}
}
