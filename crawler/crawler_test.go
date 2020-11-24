package crawler

import (
	"fmt"
	"testing"
)

type mockCrawler struct{}

// Crawl ...
func (m mockCrawler) Crawl() *Response {
	return &Response{
		Results: []Result{
			{
				Artifact: []byte{},
			},
		},
	}
}

func TestCrawler(t *testing.T) {
	m := &mockCrawler{}
	res := Start(m)

	fmt.Println(res)
}
