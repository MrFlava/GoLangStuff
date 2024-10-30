package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://postgres:secret@localhost:5432/gopostgrestest?sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	createProductTable(db)

}

func createProductTable(db *sql.DB) {
	/* Product Table
	- ID
	- Name
	- Price
	-Avaliable
	- Date Created
	*/
	query := `CREATE TABLE IF NOT EXISTS product(
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		price NUMERIC(6,2) NOT NULL,
		avaliable BOOLEAN,
		created timestamp DEFAULT NOW()
	)`

	_, err := db.Exec(query)

	if err != nil {
		log.Fatal(err)
	}
}
