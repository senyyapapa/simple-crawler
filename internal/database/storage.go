package database

type PageStorage interface {
	GetInfo(url string) *Resources
	SaveInfo(res *Resources) error
}
