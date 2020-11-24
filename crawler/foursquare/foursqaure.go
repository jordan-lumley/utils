package foursquare

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/jordan-lumley/utils/crawler"
	"github.com/peppage/foursquarego"
)

// FourSquare ...
type FourSquare struct {
	Credentials crawler.Credentials
	ZipCode     string
	Radius      int
}

// Crawl ...
func (f FourSquare) Crawl() *crawler.Response {

	// Create a new foursquare client
	client := foursquarego.NewClient(http.DefaultClient, "foursquare", f.Credentials.ClientID, f.Credentials.ClientSecret, "")

	// Get all business categories
	categories, _, err := client.Venues.Categories()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Combine all categories into one slice
	categoryIds := []string{}
	for _, categoryParent := range categories {
		categoryIds = append(categoryIds, categoryParent.ID)
	}

	// Search Venues
	venues, _, err := client.Venues.Search(&foursquarego.VenueSearchParams{
		Near:   f.ZipCode,
		Radius: f.Radius,
		Limit:  100000,
	})

	if err != nil {
		fmt.Println(err)
	}

	results := []crawler.Result{}
	for _, venue := range venues {
		venueJSON, _ := json.Marshal(venue)
		results = append(results, crawler.Result{Artifact: venueJSON})
	}

	return &crawler.Response{Results: results}
}
