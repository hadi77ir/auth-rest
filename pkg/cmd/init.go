package cmd

import (
	"auth-rest/internal/app"
	"auth-rest/internal/config"
	"auth-rest/internal/dal"
	"auth-rest/internal/db"
	"auth-rest/internal/log"
	"context"
	"github.com/hadi77ir/go-logging"
	"gorm.io/gorm"
)

func CommonInit(context context.Context, globals *app.AppGlobals, configPath string) (cfg *config.Config, logger logging.Logger, dbc *gorm.DB, repos dal.Repositories, err error) {
	// read config
	cfg, err = config.Setup(globals, configPath)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	// limit logger to level in config
	logger, err = log.Limit(globals, cfg.LogLevel)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	logger.Log(logging.DebugLevel, "done configuration")

	// setup db
	dbc, err = db.Setup(globals)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	logger.Log(logging.DebugLevel, "done setting up database")

	repos, err = dal.SetupWithGorm(globals, dbc)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	logger.Log(logging.DebugLevel, "done setting up repos")
	return
}
