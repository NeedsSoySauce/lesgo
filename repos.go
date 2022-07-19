package main

import (
	"database/sql"
	"sync"
)

type ProductsRepository interface {
	GetProducts() ([]ProductEntity, error)
	AddProduct(product Product) (*ProductEntity, error)
}

// In memory repo

type InMemoryProductsRepository struct {
	products []ProductEntity
	mutex    sync.RWMutex
}

func NewInMemoryProductsRepository(products []ProductEntity) ProductsRepository {
	return &InMemoryProductsRepository{products, sync.RWMutex{}}
}

func (repo *InMemoryProductsRepository) GetProducts() ([]ProductEntity, error) {
	repo.mutex.RLock()
	defer repo.mutex.RUnlock()
	return repo.products, nil
}

func (repo *InMemoryProductsRepository) AddProduct(product Product) (*ProductEntity, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	productEntity := ProductEntity{ID: 123, Product: product}
	repo.products = append(repo.products, productEntity)
	return &productEntity, nil
}

// Sql repo

type RemoteProductsRepository struct {
	db *sql.DB
}

func NewRemoteProductsRepository(products []ProductEntity, path *string) (ProductsRepository, error) {
	if path == nil {
		defaultPath := "default.db"
		path = &defaultPath
	}

	db, err := sql.Open("sqlite3", *path)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS some_table (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			priceInCents INTEGER
		);
	`)
	if err != nil {
		return nil, err
	}

	return &RemoteProductsRepository{db}, nil
}

func (repo *RemoteProductsRepository) GetProducts() ([]ProductEntity, error) {
	rows, err := repo.db.Query("SELECT * FROM some_table;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []ProductEntity{}

	for rows.Next() {
		var product ProductEntity
		rows.Scan(&product.ID, &product.Name, &product.PriceInCents)
		products = append(products, product)
	}

	return products, nil
}

func (repo *RemoteProductsRepository) AddProduct(product Product) (*ProductEntity, error) {
	result, err := repo.db.Exec(`
		INSERT INTO some_table (name, priceInCents)
		VALUES (?, ?);
	`, product.Name, product.PriceInCents)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &ProductEntity{uint(id), product}, nil
}
