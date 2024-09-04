package scraper

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"sync"

	"golang.org/x/net/html"
)

type Result struct {
	URL         string
	Title       string
	Description string
}

func LoadURLs(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}

	return urls, scanner.Err()
}

func ScrapeURLs(urls []string) []Result {
	var wg sync.WaitGroup
	// allocate URLs
	results := make([]Result, 0, len(urls))
	resultsChannel := make(chan Result)

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			// just logging to see concurrent behavior
			fmt.Printf("Starting to scrape: %s\n", url)
			title, description := scrape(url)
			resultsChannel <- Result{URL: url, Title: title, Description: description}
			fmt.Printf("Finished scraping: %s\n", url)
		}(url)
	}

	// closing results channel once all goroutines are done
	go func() {
		wg.Wait()
		close(resultsChannel)
	}()

	// collecting results from the channel
	for result := range resultsChannel {
		results = append(results, result)
	}

	return results
}

func scrape(url string) (string, string) {
	resp, err := http.Get(url)
	if err != nil {
		return url, "Error fetching"
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return url, "Error: Status code " + resp.Status
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return url, "Error parsing HTML"
	}

	var title, description string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" {
			title = n.FirstChild.Data
		}
		if n.Type == html.ElementNode && n.Data == "meta" {
			for _, attr := range n.Attr {
				if attr.Key == "name" && attr.Val == "description" {
					description = n.Attr[1].Val
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return title, description
}
