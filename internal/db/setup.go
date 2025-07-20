package db

import (
	"auth-rest/internal/app"
	"auth-rest/internal/config"
	"auth-rest/internal/log"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"

	"gorm.io/gorm"
)

type DBKeyType int

const DBKey DBKeyType = 0xDBDBDB

func openDB(dbType, dsn string) (gorm.Dialector, error) {
	switch dbType {
	case "mysql":
		return mysql.Open(dsn), nil
	case "pgsql":
		return postgres.Open(dsn), nil
	case "sqlite":
		return sqlite.Open(dsn), nil
	case "sqlserver":
		return sqlserver.Open(dsn), nil
	case "clickhouse":
		return clickhouse.Open(dsn), nil
	}
	return nil, fmt.Errorf("unsupported database type: %s", dbType)
}

func Setup(globals *app.AppGlobals) (*gorm.DB, error) {
	cfg := config.FromGlobals(globals)
	if cfg == nil || cfg.Database == nil {
		return nil, fmt.Errorf("config is nil")
	}

	db, err := openDB(cfg.Database.Type, cfg.Database.DSN)
	if err != nil {
		return nil, err
	}
	logger := &Logger{loggger: log.FromGlobals(globals)}
	orm, err := gorm.Open(db, &gorm.Config{Logger: logger})
	if err != nil {
		return nil, err
	}

	globals.Set(DBKey, orm)
	return orm, nil
}

func FromGlobals(globals *app.AppGlobals) *gorm.DB {
	return app.Value[*gorm.DB](globals, DBKey)
}

func FromContext(ctx fiber.Ctx) *gorm.DB {
	appCtx := app.FromContext(ctx)
	return FromGlobals(appCtx)
}

func init() {

}
