package setup

import (
	"fmt"

	env "github.com/caarlos0/env/v6"
)

func LoadEnvConfig() (EnvConfig, error) {
	envConfig := EnvConfig{}
	if err := env.Parse(&envConfig); err != nil {
		return EnvConfig{}, fmt.Errorf("failed to parse environment variables into config: %w", err)
	}

	return envConfig, nil
}

type EnvConfig struct {
	DatabaseUserName string `env:"DB_USER_NAME,required,unset"`
	DatabasePassword string `env:"DB_PASSWORD,required,unset"`
	DatabaseName     string `env:"DB_NAMES,required,unset"`
	DatabasePort     int    `env:"DB_PORT,required,unset"`
	DatabaseHost     string `env:"DB_HOST,required,unset"`
}
