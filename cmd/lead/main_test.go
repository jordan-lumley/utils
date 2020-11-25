package main

import (
	"testing"
)

type mockCrawlerClient struct{}

func (m mockCrawlerClient) Crawl() *crawler.Response {
	return &crawler.Response{
		Results: []crawler.Result{
			{
				Artifact: []byte{},
			},
		},
	}
}

func TestEngine(t *testing.T) {
	crawler := mockCrawlerClient{}
	r := crawler.Crawl()
	if r != nil {

	}
}
