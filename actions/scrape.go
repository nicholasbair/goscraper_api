package actions

import (
	"net/url"

	"github.com/gobuffalo/buffalo"
	"github.com/nicholasbair/goscraper"
)

// ScrapeHandler is a handler for scraping and stuff
func ScrapeHandler(c buffalo.Context) error {
	p := map[string][]string{}

	if m, ok := c.Params().(url.Values); ok {
		for k, v := range m {
			p[k] = v
		}
	}
	return c.Render(200, r.JSON(goscraper.Scrape(p)))
}

// GetAllScrapersHandler is a handler for returning available scrapers
func GetAllScrapersHandler(c buffalo.Context) error {
	return c.Render(200, r.JSON(goscraper.GetConfigs()))
}

// GetScraperHandler is a handler for a scraper
func GetScraperHandler(c buffalo.Context) error {
	return c.Render(200, r.JSON(goscraper.GetConfig(c.Param("provider"))))
}
