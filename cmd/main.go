package main

import (
	"log/slog"
	"main/internal/config"
	"main/internal/crawler"
	"os"
)

func main() {
	cfg := config.MustLoad()
	log := slog.New(slog.NewTextHandler(os.Stdout, nil))

	startUrl := []string{
		"https://golang.org",
		"https://google.com/",
		"https://habr.com/",
		"https://developer.mozilla.org",
		"https://yandex.ru",
	}

	crawler, err := crawler.New(log, startUrl, cfg)
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}

	crawler.Start()
}
