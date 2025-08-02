package scraper

import (
	"net/http/cookiejar"

	"github.com/gocolly/colly"
)

// SetupCollector is a function that returns a new colly collector with a default user agent
// this is used to avoid being blocked by the website
func SetupBasicCollector() *colly.Collector {
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36"
	c.AllowURLRevisit = true
	jar, err := cookiejar.New(nil)
	if err == nil {
		c.SetCookieJar(jar)
	}
	return c
}
