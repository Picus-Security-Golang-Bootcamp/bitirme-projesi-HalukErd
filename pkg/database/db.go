package db

import (
	"BasketProjectGolang/pkg/config"
	"go.uber.org/zap"
	gormPsql "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func Connect(cfg *config.Config) *gorm.DB {
	dataSourceName := cfg.DBConfig.DataSourceName
	db, err := gorm.Open(gormPsql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		zap.L().Fatal("Cannot connect to database", zap.Error(err))
	}

	origin, err := db.DB()
	if err != nil {
		zap.L().Error("Cannot get sql.DB from database", zap.Error(err))
	}

	origin.SetMaxOpenConns(cfg.DBConfig.MaxOpen)
	origin.SetMaxIdleConns(cfg.DBConfig.MaxIdle)
	origin.SetConnMaxLifetime(time.Duration(cfg.DBConfig.MaxLifetime) * time.Second)

	return db
}
