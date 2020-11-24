package main

import (
	"fmt"

	"github.com/ces/engine/pkg/crawler"
	"github.com/jordan-lumley/a1pos-utils/pkg/service"
)

const (
	clientID     = "53K5L0QOC555SUF3HLYGZUSSIE3TXS4EGOTIOKFSBPVLC3TF"
	clientSecret = "FOWCM2UIWXMXPJFTVBV5AEBUSIR1Z35VCVG1JUX5RRLKJ4XF"
)

// Crawler ...
type Crawler interface {
	Crawl()
}

func main() {
	fourSqaureCrawler := crawler.FourSquare{
		Credentials: crawler.Credentials{
			ClientID:     clientID,
			ClientSecret: clientSecret,
		},
		ZipCode: "16001",
		Radius:  5000,
	}

	fourSqaureCrawler.Crawl()

	fmt.Println("done")

	service.Run()
}
