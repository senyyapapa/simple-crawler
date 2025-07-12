package crawler

import (
	"bytes"
	"io"
	"main/internal/database"
	"main/internal/parser"
	"net/http"
	"time"
)

func (c *Crawler) worker(id int) {
	c.log.Info("worker started", "id", id)
	defer c.log.Info("worker finished", "id", id)

	for link := range c.linksChan {
		c.log.Info("Crawling link", "id", id, "url", link)
		req, err := http.NewRequest("GET", link, nil)
		if err != nil {
			c.log.Error("Failed to create request", "id", id, "url", link, "error", err)
			c.wg.Done()
			continue
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			c.log.Error("Failed to fetch", "id", id, "url", link, "error", err)
			c.wg.Done()
			continue
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			c.log.Error("Recived non-200 status code", "id", id, "url", link, "status", resp.StatusCode)
			c.wg.Done()
			continue
		}

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			c.log.Error("Failed to save resource to DB", "url", link, "error", err)
			c.wg.Done()
			continue
		}

		resources := &database.Resources{
			URL:       link,
			HTML:      string(bodyBytes),
			CrawledAt: time.Now(),
		}

		if err := c.storage.SaveInfo(resources); err != nil {
			c.log.Error("Failed to save resource to DB", "url", link, "error", err)
		}

		bodyReader := io.NopCloser(bytes.NewBuffer(bodyBytes))
		newLinks, err := parser.ExtractLinks(bodyReader, link)
		if err != nil {
			c.log.Error("Failed to parse links", "url", link, "error", err)
			c.wg.Done()
			continue
		}

		c.log.Info("Found new links", "worker", id, "count", len(newLinks), "from", link)

		if len(newLinks) > 0 {
			go func(links []string) {
				c.wg.Add(len(links))
				c.newLinksChan <- links
			}(newLinks)
		}

		c.wg.Done()

	}
}
