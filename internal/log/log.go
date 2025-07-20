package log

import (
	"github.com/hadi77ir/go-env"
	"os"

	"github.com/hadi77ir/go-logging"
	jlog "github.com/hadi77ir/go-logging/json"

	"auth-rest/internal/app"
)

const DefaultLogLevel = "${LOG_LEVEL:-info}"

type LoggerKeyType int

const LoggerKey LoggerKeyType = 0xDEADBEEF

func Setup(globals *app.AppGlobals) logging.Logger {
	var logger logging.Logger
	logger = jlog.New(os.Stderr)
	globals.Set(LoggerKey, logger)
	return logger
}

func Limit(globals *app.AppGlobals, level string) (logging.Logger, error) {
	if level == "" {
		level, _ = env.ExpandEnv(DefaultLogLevel)
	}
	parsedLevel, err := logging.ParseLevel(level)
	if err != nil {
		return nil, err
	}
	var logger logging.Logger
	logger = FromGlobals(globals)
	if logger == nil {
		logger = logging.NoOpLogger(0)
	}
	logger = logging.Limit(logger, parsedLevel)
	globals.Set(LoggerKey, logger)
	return logger, nil
}

func FromGlobals(globals *app.AppGlobals) logging.Logger {
	return app.Value[logging.Logger](globals, LoggerKey)
}
