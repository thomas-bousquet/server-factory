package logger

import (
	"encoding/json"

	"github.com/thomas-bousquet/server-factory/pkg/server/config"
	"go.uber.org/zap"
)

func NewLogger(c config.Config) (*zap.Logger, error) {
	encoding := "json"

	if c.Env == "dev" {
		encoding = "console"
	}

	rawJSON := []byte(`{
	  "level": "` + c.LogLevel + `",
	  "encoding": "` + encoding + `",
	  "outputPaths": ["stdout"],
	  "errorOutputPaths": ["stderr"],
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase"
	  }
	}`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		return nil, err
	}

	return cfg.Build()
}
