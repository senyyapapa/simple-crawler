package crawler

import (
	"log/slog"
	"main/internal/client"
	"main/internal/config"
	"main/internal/database"
	"main/internal/database/sql"
	"sync"
	"time"
)

type Crawler struct {
	storage    database.PageStorage
	log        *slog.Logger
	startUrl   []string
	cfg        *config.Config
	httpClient client.Fetcher
	wg         sync.WaitGroup
	linksChan  chan string
	visited    map[string]bool
	mu         sync.Mutex
}

//TODO: Добавить Graceful Shutdown

func New(log *slog.Logger, startUrl []string, cfg *config.Config) (*Crawler, error) {
	dialector, config := sql.InitGorm(cfg, log)
	db, err := database.NewSQLStorage(dialector, config)
	if err != nil {
		return nil, err
	}

	httpClient := client.New(10 * time.Second)

	return &Crawler{
		storage:    db,
		log:        log,
		startUrl:   startUrl,
		cfg:        cfg,
		httpClient: httpClient,
		linksChan:  make(chan string, 1000),
		visited:    make(map[string]bool),
	}, nil
}

func (c *Crawler) Start() {
	c.log.Info("Starting crawler...")

	for i := 0; i < c.cfg.WorkersCount; i++ {
		go c.worker(i)
	}

	for _, u := range c.startUrl {
		c.addToQueue(u)
	}

	c.wg.Wait()

	close(c.linksChan)

	c.log.Info("Crawler finished")
}
