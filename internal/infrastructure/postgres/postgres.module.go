package postgres

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/dangduoc08/ginject-cms-api/internal/infrastructure/config"
	"github.com/dangduoc08/ginject/common"
	"github.com/dangduoc08/ginject/core"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type GormLogger struct {
	Logger common.Logger
}

func (instance *GormLogger) LogMode(logLevel gormLogger.LogLevel) gormLogger.Interface {
	return instance
}

func (instance *GormLogger) Info(c context.Context, msg string, data ...any) {
	instance.Logger.Info(msg, data)
}

func (instance *GormLogger) Warn(c context.Context, msg string, data ...any) {
	instance.Logger.Warn(msg, data)
}

func (instance *GormLogger) Error(c context.Context, msg string, data ...any) {
	instance.Logger.Error(msg, data)
}

func (instance *GormLogger) Trace(c context.Context, begin time.Time, cb func() (sql string, rowsAffected int64), err error) {
	sql, rowsAffected := cb()
	sql = regexp.MustCompile(`\s+`).ReplaceAllString(sql, " ")
	sql = strings.TrimSpace(sql)

	if err != nil {
		instance.Logger.Error(err.Error(), "sql", sql, "rowsAffected", rowsAffected)
	} else {
		instance.Logger.Debug("GORM", "sql", sql, "rowsAffected", rowsAffected)
	}
}

var PostgresModule = func(
	logger common.Logger,
) *core.Module {

	sslmode := "disable"
	if config.Env.PostgresSSL {
		sslmode = "require"
	}

	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=Asia/Shanghai",
		config.Env.PostgresHost,
		config.Env.PostgresUser,
		config.Env.PostgresPassword,
		config.Env.PostgresDB,
		config.Env.PostgresPort,
		sslmode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: &GormLogger{logger},
		NamingStrategy: schema.NamingStrategy{
			NameReplacer: strings.NewReplacer("Model", ""),
		},
	})

	if err != nil {
		logger.Fatal("PostgresSQL", "error", err.Error(), "connected", false)
	} else {
		logger.Info("PostgresSQL", "connected", true)
	}

	provider := PostgresProvider{
		DB: db,
	}

	module := core.ModuleBuilder().
		Providers(provider).
		Build()

	module.IsGlobal = true

	return module
}
