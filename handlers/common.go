package handlers

import (
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"os"
)



func NewCollyCollector() *colly.Collector {
	c := colly.NewCollector(
		colly.UserAgent(os.Getenv("User-Agent")),
	)

	extensions.RandomUserAgent(c)

	return c
}