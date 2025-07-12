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
		"https://github.com",
		"habr.com/ru/companies/domclick/articles/592087/",
		"https://pkg.go.dev/golang.org/x/net/html",
		"https://monkeytype.com/",
		"https://comx.life/10072-chernyj-klever.html#chapters",
	}

	crawler, err := crawler.New(log, startUrl, cfg)
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}

	crawler.Start()
}
