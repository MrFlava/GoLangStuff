package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Product struct {
	Name      string
	Price     float64
	Avaliable bool
}

func main() {
	connStr := "postgres://postgres:secret@localhost:5433/gopostgrestest?sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	product := Product{"book", 15.55, true}
	pk := insertProduct(db, product)

	fmt.Printf("ID = %d\n", pk)

}

// func createProductTable(db *sql.DB) {
// 	/* Product Table
// 	- ID
// 	- Name
// 	- Price
// 	-Avaliable
// 	- Date Created
// 	*/
// 	query := `CREATE TABLE IF NOT EXISTS product(
// 		id SERIAL PRIMARY KEY,
// 		name VARCHAR(100) NOT NULL,
// 		price NUMERIC(6,2) NOT NULL,
// 		avaliable BOOLEAN,
// 		created timestamp DEFAULT NOW()
// 	)`

// 	_, err := db.Exec(query)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func insertProduct(db *sql.DB, product Product) int {
	query := `INSERT INTO product(name, price, avaliable)
		VALUES ($1, $2, $3) RETURNING id`

	var pk int
	err := db.QueryRow(query, product.Name, product.Price, product.Avaliable).Scan(&pk)

	if err != nil {
		log.Fatal(err)
	}

	return pk
}
