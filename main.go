package main

import (
	"concurrent-web-scraper/scraper"
	"fmt"
	"log"
)

func main() {
	urls, err := scraper.LoadURLs("data/urls.txt")
	if err != nil {
		log.Fatalf("Error loading URLs: %v", err)
	}

	results := scraper.ScrapeURLs(urls)

	for _, result := range results {
		fmt.Printf("URL: %s\nTitle: %s\nDescription: %s\n\n", result.URL, result.Title, result.Description)
	}
}
