package database

import "gorm.io/gorm"

type SQLStorage struct {
	DB *gorm.DB
}

func NewSQLStorage(dialector gorm.Dialector, config *gorm.Config) (*SQLStorage, error) {
	db, err := gorm.Open(dialector, config)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Resources{})

	return &SQLStorage{DB: db}, nil
}

func (db *SQLStorage) GetInfo(url string) *Resources {
	var res Resources
	db.DB.Where("url = ?", url).First(&res)
	return &res
}

func (db *SQLStorage) SaveInfo(res *Resources) error {
	result := db.DB.Create(res)
	return result.Error
}
