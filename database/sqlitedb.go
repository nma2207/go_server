package database

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

//provider to SQLite database
type SqliteDB struct {
	DB *sql.DB
}

func (db *SqliteDB) InitDatabase() error {
	var err error
	db.DB, err = sql.Open("sqlite3", "products.db")
	return err
}

func (db *SqliteDB) AddProduct(product Product) (uint64, error) {
	row := db.DB.QueryRow("SELECT * FROM products WHERE name = $1",
		product.Name)

	if row.Scan() != sql.ErrNoRows {
		return 0, errors.New("Has same name")
	}

	result, err := db.DB.Exec(`INSERT INTO products (name, cost) 
                        VALUES ($1, $2)`, product.Name, product.Cost)

	if err != nil {
		return 0, err
	}
	id, _ := result.LastInsertId()
	return uint64(id), nil
}

func (db *SqliteDB) DeleteProduct(id uint64) error {
	_, err := db.DB.Exec("DELETE FROM products WHERE id = $1", id)
	return err
}

func (db *SqliteDB) Change(product Product) error {
	_, err := db.DB.Exec(`UPDATE products 
                        SET name = $1, cost = $2 
                        WHERE id = $3`, product.Name, product.Cost, product.Id)
	return err
}

func (db *SqliteDB) Get(id uint64) (Product, error) {
	row := db.DB.QueryRow("SELECT * FROM products WHERE id = $1", id)
	product := Product{}
	err := row.Scan(&product.Id, &product.Name, &product.Cost)

	return product, err
}

func (db *SqliteDB) GetAll() ([]Product, error) {
	rows, err := db.DB.Query("SELECT * FROM products ORDER BY cost")
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

func (db *SqliteDB) Close() {
	fmt.Println("Close db")
	db.DB.Close()
}
