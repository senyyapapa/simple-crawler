package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log/slog"
	sl "main/libs/logger"
)

type SQLiteStorage struct {
	DB  *gorm.DB
	log *slog.Logger
}

func NewSQLiteStorage(db_url string, log *slog.Logger) (*SQLiteStorage, error) {
	db, err := gorm.Open(sqlite.Open(db_url), &gorm.Config{})
	if err != nil {
		log.Error("Connection failed to database", sl.Err(err))
		return nil, err
	}

	err = db.AutoMigrate(&Resources{})
	if err != nil {
		log.Error("Migration failed", sl.Err(err))
		return nil, err
	}

	return &SQLiteStorage{
		DB:  db,
		log: log,
	}, nil
}

func (s *SQLiteStorage) GetInfo(url string) *Resources {
	var res Resources
	s.DB.Where("url = ?", url).First(&res)
	return &res
}
