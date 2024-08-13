package scraper

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)



type Scraper interface {
	getTeamx(c *colly.Collector)
}

type WebsiteAScraper struct{}

func Run() {
	fmt.Println("Running Scraper...")

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	Scraper.getTeamx(c)

}

func (w *WebsiteAScraper) getTeamx(c *colly.Collector) {
	totalPages := 0
	visitedPages := map[string]bool{}
	var linksSlice []string

	c.OnHTML(".last-chapter .pagination", func(e *colly.HTMLElement) {
		pageLinks := e.DOM.Find(".page-item a")
		pageLinks.Each(func(_ int, s *goquery.Selection) {
			// Extract the href attribute of each pagination link
			href, _ := s.Attr("href")
			if href != "" {
				absURL := e.Request.AbsoluteURL(href)
				if !visitedPages[absURL] {
					visitedPages[absURL] = true
					totalPages++ // Increment total pages count
				}
			}
		})
		// fmt.Printf("Total pages: %d\n", totalPages)
	})

	// Function to handle the main content scraping
	c.OnHTML(".listupd .bs", func(e *colly.HTMLElement) {
		e.ForEach(".bsx a", func(_ int, child *colly.HTMLElement) {
			link := child.Attr("href")
			// fmt.Printf("Link found: %s\n", link)
			linksSlice = append(linksSlice, link)
		})
	})

	// Function to handle pagination and visit pages
	c.OnHTML(".last-chapter .pagination", func(e *colly.HTMLElement) {
		var nextPage string
		e.ForEach(".page-item a", func(_ int, child *colly.HTMLElement) {
			if child.Attr("rel") == "next" && child.ChildAttr(".page-item .disabled", "aria-disabled") != "true" {
				nextPage = child.Attr("href")
			}
		})

		if nextPage != "" {
			nextPageURL := e.Request.AbsoluteURL(nextPage)
			// fmt.Println("Found next page:", nextPageURL)
			c.Visit(nextPageURL)
		} else {
			fmt.Println("No more pages to visit.")
		}
	})

	// Start the scraping process with the initial URL
	startURL := "https://teamoney.site/series/"
	c.Visit(startURL)

	if len(linksSlice) != 0 {
		fmt.Println("Data: ", len(linksSlice), "total Pages: ", totalPages)
		// fmt.Println(len(linksSlice))
	}
}

// func getAzora() {

// }
