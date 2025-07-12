package database

import "time"

type Resources struct {
	URL  string `gorm:"primaryKey"`
	HTML string
	//TODO: Add parsing of css and js. Add language of page and meta-data
	CrawledAt time.Time
}
