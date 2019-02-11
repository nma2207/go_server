package database

type Product struct {
	Id   uint64  `json="id"`
	Name string  `json="name"`
	Cost float64 `json="cost"`
}

type IDataBase interface {
	AddProduct(product Product) (uint64, error)
	DeleteProduct(id uint64) error
	Change(product Product) error
	Get(id uint64) (Product, error)
	GetAll() ([]Product, error)
}

type DataBase struct {
	Products []Product
}

func (db *DataBase) AddProduct(product Product) (uint64, error) {
	db.Products = append(db.Products, product)
	return 0, nil
}

func (db *DataBase) DeleteProduct(id uint64) error {
	var newProducts []Product
	for _, p := range db.Products {
		if p.Id != id {
			newProducts = append(newProducts, p)
		}
	}

	db.Products = newProducts
	return nil
}

func (db *DataBase) Change(product Product) error {
	for _, p := range db.Products {
		if p.Id == product.Id {
			p.Name = product.Name
			p.Cost = product.Cost
		}
	}
	return nil
}

func (db *DataBase) Get(id uint64) (Product, error) {
	for _, p := range db.Products {
		if p.Id == id {
			return p, nil
		}
	}

	return db.Products[0], nil
}

func (db *DataBase) GetAll() ([]Product, error) {
	return db.Products, nil
}
