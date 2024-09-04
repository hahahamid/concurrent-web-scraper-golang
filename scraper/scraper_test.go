package scraper

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoadURLs(t *testing.T) {
	urls, err := LoadURLs("../data/urls.txt")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(urls) == 0 {
		t.Fatal("Expected URLs, got none")
	}
}

func TestScrape(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<html><head><title>Mock Title</title><meta name="description" content="Mock Description"></head><body></body></html>`))
	}))
	defer server.Close()

	title, description := scrape(server.URL)
	if title == "" {
		t.Fatal("Expected a title, got none")
	}
	if description == "" {
		t.Fatal("Expected a description, got none")
	}
}
