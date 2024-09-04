# Concurrent Web Scraper

This project demonstrates concurrency in Golang by implementing a web scraper that fetches data from multiple URLs concurrently. It uses goroutines and channels to fetch and parse data from multiple websites simultaneously, showcasing how Golang can handle multiple tasks efficiently.

## Technologies Used
- Go (Golang)
- HTML parsing with golang.org/x/net/html
- Concurrency with goroutines and channels


## Features
- Concurrently fetches and parses HTML from multiple websites.
- Extracts titles and meta descriptions.
- Displays results in a structured format.

## Setup
1. Clone the repository.
2. Run `go mod tidy` to install dependencies.
3. Add URLs to `data/urls.txt`.
4. Run the application with `go run main.go`.

## Testing
Run tests with `go test ./scraper`.