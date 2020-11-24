package crawler

import (
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/peppage/foursquarego"
)

var (
	wg sync.WaitGroup
)

// Credentials ...
type Credentials struct {
	ClientID     string
	ClientSecret string
}

// FourSquare ...
type FourSquare struct {
	Credentials
	ZipCode string
	Radius  int
}

// Crawl ...
func (f FourSquare) Crawl() {

	// Create a new foursquare client
	client := foursquarego.NewClient(http.DefaultClient, "foursquare", f.ClientID, f.ClientSecret, "")

	// Get all business categories
	categories, _, err := client.Venues.Categories()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Combine all categories into one slice
	categoryIds := []string{}
	for _, category := range categories {
		categoryIds = append(categoryIds, category.ID)
	}

	// Search Venues
	venues, _, err := client.Venues.Search(&foursquarego.VenueSearchParams{
		Near:       f.ZipCode,
		Radius:     f.Radius,
		CategoryID: categoryIds,
	})

	if err != nil {
		fmt.Println(err)
	}

	_ = venues
	fmt.Println(len(venues))

	// prettyJSON, err := json.MarshalIndent(venues, "", "    ")

	// fmt.Printf("%s\n", string(prettyJSON))
}
