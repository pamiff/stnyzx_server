package config

import (
	"errors"
	"os"

	"github.com/micro/go-micro/config/encoder/toml"
	"github.com/micro/go-micro/config/source"

	"github.com/micro/go-micro/config/source/file"
	"github.com/micro/go-plugins/config/source/configmap"

	"github.com/micro/go-micro/config"
)

var (
	ProgramEnvironment string
)

const (
	ServiceName = "auth"
)

func init() {
	ProgramEnvironment = os.Getenv("program_environment")
}

func LoadConfig() (cfg Config, err error) {
	c := config.NewConfig()
	enc := toml.NewEncoder()
	if ProgramEnvironment == "dev" {
		if err = c.Load(
			file.NewSource(
				file.WithPath("../config/config.toml"),
				source.WithEncoder(enc),
			),
		); err != nil {
			return
		}
		if err = c.Scan(&cfg); err != nil {
			return
		}
		return
		// return config.LoadFile("../config/config.toml")
		// }
	} else if ProgramEnvironment == "test" {
		if err = c.Load(
			configmap.NewSource(
				configmap.WithNamespace("stnyzx-test"),
				configmap.WithName(ServiceName),
				source.WithEncoder(enc),
			),
		); err != nil {
			return
		}
		if err = c.Scan(&cfg); err != nil {
			return
		}
		return
		// return config.Load(
		// 	configmap.NewSource(
		// 		configmap.WithNamespace("stnyzx-test"),
		// 		configmap.WithName("auth"),
		// 	),
		// )
	} else if ProgramEnvironment == "prod" {
		if err = c.Load(
			configmap.NewSource(
				configmap.WithNamespace("stnyzx-prod"),
				configmap.WithName(ServiceName),
				source.WithEncoder(enc),
			),
		); err != nil {
			return
		}
		if err = c.Scan(&cfg); err != nil {
			return
		}
		return
		// return config.Load(
		// 	configmap.NewSource(
		// 		configmap.WithNamespace("stnyzx-prod"),
		// 		configmap.WithName("auth"),
		// 	),
		// )
	}
	err = errors.New("env var $program_environment must be in [dev, test, prod]")
	return
}
