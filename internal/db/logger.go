package db

import (
	"context"
	"github.com/hadi77ir/go-logging"
	"gorm.io/gorm/logger"
	"time"
)

type Logger struct {
	loggger logging.Logger
}

func (l *Logger) LogMode(level logger.LogLevel) logger.Interface {
	return l
}

func (l *Logger) log(level logging.Level, s string, i ...interface{}) {
	args := []any{s}
	args = append(args, i...)
	l.loggger.Log(level, args...)
}
func (l *Logger) Info(ctx context.Context, s string, i ...interface{}) {
	l.log(logging.InfoLevel, s, i...)
}

func (l *Logger) Warn(ctx context.Context, s string, i ...interface{}) {
	l.log(logging.WarnLevel, s, i...)
}

func (l *Logger) Error(ctx context.Context, s string, i ...interface{}) {
	l.log(logging.ErrorLevel, s, i...)
}

func (l *Logger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, rowsAffected := fc()
	l.loggger.Log(logging.TraceLevel, "query began at ", begin.Format(time.RFC3339), ", with ", rowsAffected, " rows affected: ", sql, ", error: ", err)
}

var _ logger.Interface = &Logger{}
