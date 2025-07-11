package sql

import (
	slogGorm "github.com/orandin/slog-gorm"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log/slog"
	"main/internal/config"
	"time"
)

func InitGorm(cfg *config.Config, log *slog.Logger) (gorm.Dialector, *gorm.Config) {
	dbOptions := sqlite.New(sqlite.Config{
		DSN: cfg.DbUrl,
	},
	)
	gormSlogOpt := append(
		[]slogGorm.Option{},
		slogGorm.WithHandler(log.Handler()),
		slogGorm.WithSlowThreshold(time.Millisecond*200),
		slogGorm.SetLogLevel(slogGorm.DefaultLogType, slog.LevelDebug),
	)

	if cfg.Env == "prod" {
		gormSlogOpt = append(
			gormSlogOpt,
			slogGorm.WithIgnoreTrace(),
		)
	} else if *cfg.LoggerLevel == slog.LevelDebug {
		gormSlogOpt = append(
			gormSlogOpt,
			slogGorm.WithTraceAll(),
		)
	}
	gormCfg := &gorm.Config{
		Logger:               slogGorm.New(gormSlogOpt...),
		PrepareStmt:          true,
		DisableAutomaticPing: true,
	}

	return dbOptions, gormCfg

}
