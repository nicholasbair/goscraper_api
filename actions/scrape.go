package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/nicholasbair/goscraper"
)

// ScrapeHandler is a handler for scraping and stuff
func ScrapeHandler(c buffalo.Context) error {
	return c.Render(200, r.JSON(goscraper.Scrape()))
}

// GetAllScrapersHandler is a handler for returning available scrapers
func GetAllScrapersHandler(c buffalo.Context) error {
	return c.Render(200, r.JSON(goscraper.Scrape()))
}

// TODO
// Fix url, just returning base url
// Setup routes for each scraper, need to handle query params
