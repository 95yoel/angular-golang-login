package dbs

import (
	dbconfig "backend/src/database_config"
	"database/sql"
	"fmt"
	"log"
)

// global variable db
var DB *sql.DB

/*
Open database
*/

func OpenDatabase(connStr string) (*sql.DB, error) {
	return sql.Open("postgres", connStr)
}

/*
Check if the database Works
*/

func PingDatabase(db *sql.DB) error {
	return db.Ping()
}

/*
Initializes the PostgreSQL database connection.
*/

func InitializeDb() {

	// Get connection string
	connStr := dbconfig.GetConnectionString()

	// Open the database with connection string
	db, err := OpenDatabase(connStr)
	handleError(err)

	// Check if database is working
	err = PingDatabase(db)
	handleError(err)

	// Set the global DB variable to the opened database
	DB = db

	fmt.Println("PostgreSQL connected succesfully")

}

// Close database

func CloseDb() {
	if DB != nil {
		DB.Close()
		fmt.Println("Database connection closed.")
	}
}

// Handle errors

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
