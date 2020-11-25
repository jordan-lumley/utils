package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

const (
	host     = "54.196.228.234"
	port     = 5432
	user     = "atlas_crm_api"
	password = "monstrous cymbal caucasian mud"
	dbname   = "atlascrm"
)

// NewConnection ...
func NewConnection() *sql.DB {
	connInfo := fmt.Sprintf("host=%s port=%d user=%s password='%s' dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connInfo)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return db
}
