package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AppName                  string        `envconfig:"X_APP_NAME" required:"true"`
	AppVersion               string        `envconfig:"X_APP_VERSION" required:"true"`
	Env                      string        `envconfig:"X_ENV" required:"true"`
	AppPort                  int           `envconfig:"X_APP_PORT" default:"8080"`
	LogLevel                 string        `envconfig:"X_LOG_LEVEL" default:"info"`
	ReadTimeout              time.Duration `envconfig:"X_READ_TIMEOUT" default:"5s"`
	WriteTimeout             time.Duration `envconfig:"X_WRITE_TIMEOUT" default:"5s"`
	IdleTimeout              time.Duration `envconfig:"X_IDLE_TIMEOUT" default:"5s"`
	GracefullShutdownTimeout time.Duration `envconfig:"X_GRACEFULL_SHUTDOWN_TIMEOUT" default:"10s"`
}

func NewConfig() (Config, error) {
	c := Config{}
	err := envconfig.Process("", &c)

	if err != nil {
		return Config{}, err
	}

	return c, nil
}

func NewTestConfig() Config {
	return Config{
		AppName:    "test-app-name",
		AppVersion: "0.0.0",
		Env:        "dev",
	}
}
