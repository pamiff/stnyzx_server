package config_test

import (
	//"github.com/micro/go-micro/config"
	"fmt"
	"testing"
	"github.com/pamiff/stnyzx_server/auth/config"
)

func TestLoadConfig(t *testing.T) {
	cfg, err := config.LoadConfig()
	fmt.Println(cfg, err)
}