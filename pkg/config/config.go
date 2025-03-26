package config

import (
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

func LoadEnv(envName ...string) (cfg Cfg, err error) {
	envFiles := make([]string, 0, len(envName))
	for _, v := range envName {
		if _, err := os.Stat(v); err == nil {
			envFiles = append(envFiles, v)
		}
	}
	if len(envFiles) != 0 {
		err = godotenv.Load(envFiles...)
		if err != nil {
			return
		}
	}
	cfg, err = env.ParseAs[Cfg]()
	return
}
