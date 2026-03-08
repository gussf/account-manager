package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	HTTPServerPort           int           `mapstructure:"http_server_port"`
	HTTPServerRequestTimeout time.Duration `mapstructure:"http_server_request_timeout"`
	DatabaseHostname         string        `mapstructure:"database_hostname"`
	DatabasePort             int           `mapstructure:"database_port"`
	DatabaseName             string        `mapstructure:"database_name"`
	DatabaseUser             string        `mapstructure:"database_user"`
	DatabasePassword         string        `mapstructure:"database_password"`
}

func Load() (*Config, error) {
	viper.SetConfigName("local.config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("fatal error config file: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
