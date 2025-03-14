package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type parseConfig struct {
	usersHost   string `env:"USERS_HOST" env-default:"localhost"`
	usersPort   string `env:"USERS_PORT" env-default:"50000"`
	timeHost    string `env:"TIME_HOST" env-default:"localhost"`
	timePort    string `env:"TIME_PORT" env-default:"50000"`
	gatewayHost string `env:"HOST" env-default:"locahost"`
	gatewayPort string `env:"PORT" env-default:"1231"`
}

type Config struct {
	UsersUrlService string
	TimeUrlService  string
	GatewayURL      string
}

func MustLoad() (*Config, error) {
	var cfg parseConfig

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, fmt.Errorf("MustLoad: failed read env file and parse to cfg struct: %w", err)
	}

	fields := map[string]string{"gatewayHost": cfg.gatewayHost,
		"usersHost":   cfg.usersHost,
		"timeHost":    cfg.timeHost,
		"gatewayPort": cfg.gatewayPort,
		"usersPort":   cfg.usersPort,
		"timePort":    cfg.timePort,
	}

	var emptyFields []string
	for key, value := range fields {
		if value == "" {
			emptyFields = append(emptyFields, key)
		}
	}

	if len(emptyFields) > 0 {
		return nil, fmt.Errorf("MustLoad: empty fields: %v", emptyFields)
	}
	usersURL := fmt.Sprintf("%s:%s", cfg.usersHost, cfg.usersPort)
	timeURL := fmt.Sprintf("%s:%s", cfg.timeHost, cfg.timePort)
	gatewayURL := fmt.Sprintf("%s:%s", cfg.gatewayHost, cfg.gatewayPort)

	return &Config{
		UsersUrlService: usersURL,
		TimeUrlService:  timeURL,
		GatewayURL:      gatewayURL,
	}, nil
}
