package main

import (
	"database/sql"
	"fmt"
	"github.com/tlsarchiver/dbconnector"
)

var (
	db *sql.DB
)

func main() {
	fmt.Println("TLS Archiver - Analyzer")

	// Retrieve the configuration from the environment
	dbConfig := dbconnector.ParseConfiguration()

	// Setup the DB
	db = dbconnector.SetupDB(dbConfig)

	// Look for duplicate public keys
	findDuplicates()
}
