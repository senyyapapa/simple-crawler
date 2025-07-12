package crawler

func (c *Crawler) queueManager() {
	for links := range c.newLinksChan {
		for _, link := range links {
			c.mu.Lock()
			if !c.visited[link] {
				c.visited[link] = true
				c.mu.Unlock()
				c.linksChan <- link
			} else {
				c.mu.Unlock()
				c.wg.Done()
			}

		}
	}
}
