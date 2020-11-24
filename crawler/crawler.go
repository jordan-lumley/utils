package crawler

import (
	"sync"
)

var (
	wg sync.WaitGroup
)

// Client ...
type Client interface {
	Crawl() *Response
}

// Response ...
type Response struct {
	Results []Result
}

// Result ...
type Result struct {
	Artifact []byte
}

// Credentials ...
type Credentials struct {
	ClientID     string
	ClientSecret string
}

// Start ...
func Start(crawler Client) *Response {
	return crawler.Crawl()
}
