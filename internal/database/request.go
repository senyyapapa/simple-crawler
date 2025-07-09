package database

type Resources struct {
	URL  string `gorm:"primaryKey"`
	HTML string
}
