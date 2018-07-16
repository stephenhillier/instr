package database

import (
	"log"

	"github.com/jmoiron/sqlx"

	// register postgres driver
	_ "github.com/lib/pq"
)

// DB is the database connection opened by InitDB
var DB *sqlx.DB

// InitDB initializes the database connection
func InitDB(connectionConfig string) {
	var err error
	DB, err = sqlx.Open("postgres", connectionConfig)
	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Database connection initialized")
}

/*
Thermistor reading database reads/writes
*/

// StoreReading adds a reading into the database
func StoreReading(resistance float32, device string) error {
	reading := `INSERT INTO reading (device, resistance) VALUES ($1, $2)`

	_, err := DB.Exec(reading, device, resistance)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
