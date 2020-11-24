package foursquare

import (
	"fmt"
	"testing"

	"github.com/jordan-lumley/utils/crawler"
)

type mockFoursquare struct{}

// Crawl ...
func (m mockFoursquare) Crawl() *crawler.Response {
	return &crawler.Response{
		Results: []crawler.Result{
			{
				Artifact: []byte{},
			},
		},
	}
}

func TestCrawler(t *testing.T) {
	m := &mockFoursquare{}
	res := m.Crawl()

	fmt.Println(res)
}
