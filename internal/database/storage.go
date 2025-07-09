package database

type CrawlerStorage interface {
	GetInfo(url string) *Resources
}
