package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	Id   uint64
	Name string
	Cost float64
}

func add(db *sql.DB, p Product) (uint64, error) {
	rows, err := db.Query("SELECT * FROM products WHERE name = $1",
		p.Name)
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		return 0, errors.New("Has same name")
	}

	result, err := db.Exec(`INSERT INTO products (name, cost) 
							VALUES ($1, $2)`, p.Name, p.Cost)

	if err != nil {
		panic(err)
	}
	id, _ := result.LastInsertId()
	return uint64(id), nil
}

func delete(db *sql.DB, id uint64) error {
	_, err := db.Exec("DELETE FROM products WHERE id = $1", id)
	return err
}

func get(db *sql.DB, id uint64) (Product, error) {
	row := db.QueryRow("SELECT * FROM products WHERE id = $1", id)
	product := Product{}
	err := row.Scan(&product.Id, &product.Name, &product.Cost)

	return product, err
}

func getAllSort(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("SELECT * FROM products ORDER BY name")

	if err != nil {
		return make([]Product, 0), err
	}

	var products []Product

	for rows.Next() {
		p := Product{}
		err = rows.Scan(&p.Id, &p.Name, &p.Cost)
		if err != nil {
			return make([]Product, 0), err
		}
		products = append(products, p)
	}
	return products, nil
}

func main() {
	db, err := sql.Open("sqlite3", "products.db")
	if err != nil {
		fmt.Println("Errr!")
		panic(err)
	}
	defer db.Close()
	id, err := add(db, Product{0, "AIphone", 100.0})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(id)

	p, err := get(db, 2)
	fmt.Println(p, err)
	products, err := getAllSort(db)

	fmt.Println(products, err)

}
