package database

import "time"

type Resources struct {
	URL       string `gorm:"primaryKey"`
	HTML      string
	CrawledAt time.Time
}
