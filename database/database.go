package database

type Product struct {
	Id   uint64  `json:"id"`
	Name string  `json:"name"`
	Cost float64 `json:"cost"`
}

type IDataBase interface {
	InitDatabase() error
	AddProduct(product Product) (uint64, error)
	DeleteProduct(id uint64) error
	Change(product Product) error
	Get(id uint64) (Product, error)
	GetAll() ([]Product, error)
	Close()
}

type StubDataBase struct {
	Products []Product
}

func (db *StubDataBase) InitDatabase() error {
	return nil
}

func (db *StubDataBase) AddProduct(product Product) (uint64, error) {
	product.Id = uint64(len(db.Products))

	db.Products = append(db.Products, product)
	return uint64(len(db.Products) - 1), nil
}

func (db *StubDataBase) DeleteProduct(id uint64) error {
	var newProducts []Product
	for _, p := range db.Products {
		if p.Id != id {
			newProducts = append(newProducts, p)
		}
	}

	db.Products = newProducts
	return nil
}

func (db *StubDataBase) Change(product Product) error {
	for _, p := range db.Products {
		if p.Id == product.Id {
			p.Name = product.Name
			p.Cost = product.Cost
		}
	}
	return nil
}

func (db *StubDataBase) Get(id uint64) (Product, error) {
	for _, p := range db.Products {
		if p.Id == id {
			return p, nil
		}
	}

	return Product{}, nil
}

func (db *StubDataBase) GetAll() ([]Product, error) {
	return db.Products, nil
}

func (db *StubDataBase) Close() {

}
