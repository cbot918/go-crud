package configs

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Name string
		Port string
		Host string
	}
	Connections struct {
		MySQL struct {
			Host     string
			DB       string
			User     string
			Password string
			// MaxOpen  int
			// MaxIdle  int
		}
	}
}

func New() (*Config, error) {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName("app")
	v.AddConfigPath("./configs")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config error: %s", err)
	}
	cfg := &Config{}
	err := v.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("marshal error: %s", err)
	}
	return cfg, nil
}
