package main

import (
	"database/sql"
	"fmt"

	cesDb "github.com/ces/engine/pkg/db"
	"github.com/jordan-lumley/utils/crawler"
	"github.com/jordan-lumley/utils/crawler/foursquare"
	"github.com/jordan-lumley/utils/service"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

const (
	clientID     = "53K5L0QOC555SUF3HLYGZUSSIE3TXS4EGOTIOKFSBPVLC3TF"
	clientSecret = "FOWCM2UIWXMXPJFTVBV5AEBUSIR1Z35VCVG1JUX5RRLKJ4XF"
)

func main() {
	db = cesDb.NewConnection()

	defer db.Close()

	zipCodes := []string{"16001", "16602", "16025", "16027", "16033", "16041", "16029", "16053"}
	for _, zipCode := range zipCodes {
		fourSqaureCrawler := foursquare.FourSquare{
			Credentials: crawler.Credentials{
				ClientID:     clientID,
				ClientSecret: clientSecret,
			},
			ZipCode: zipCode,
			Radius:  100000,
		}

		response := crawler.Start(fourSqaureCrawler)

		fmt.Println(response)
	}

	service.Run()
}

// func crawlDb(results []crawler.Result) {
// 	var wg sync.WaitGroup

// 	for _, result := range results {
// 		wg.Add(1)
// 		go func(result crawler.Result) {
// 			defer wg.Done()

// 			var venue foursquarego.Venue
// 			err := json.Unmarshal(result.Artifact, &venue)
// 			if err != nil {
// 				fmt.Println(err)
// 				os.Exit(1)
// 			}

// 			if venue.ID == "4c792177566db60c51b7430e" {
// 				fmt.Println("FOUND IT")
// 			}
// 			// fmt.Println(venue.Location.Address)

// 			address := venue.Location.FormattedAddress[0]
// 			if address == "524 32nd St (at Crescent Rd.)" {
// 				fmt.Println("FOUND ADDRESS")
// 			}

// 			// if venue.Location.City == "Evans City" {
// 			// 	println("addhaa")
// 			// }

// 			// fmt.Printf("SELECT UPPER(document ->> 'city') FROM lead where UPPER(document ->> 'city') similar to upper(%s)", "'%("+venue.Location.City+")%'")

// 			var lead models.Lead
// 			rows, err := db.Query("SELECT lead, UPPER(document ->> 'city') as DocSearch FROM lead where UPPER(document ->> 'city') similar to upper($1)", "'%("+venue.Location.City+")%'")
// 			if err == sql.ErrNoRows {
// 				return
// 			}

// 			defer rows.Close()

// 			if err != nil {
// 				fmt.Println(err)
// 				os.Exit(1)
// 			}

// 			for rows.Next() {
// 				if err := rows.Scan(&lead); err != nil {
// 					// Check for a scan error.
// 					// Query rows will be closed with defer.
// 					log.Fatal(err)
// 				}

// 				fmt.Println(lead)
// 			}

// 		}(result)
// 	}

// 	wg.Wait()

// 	fmt.Println("done")
// }
