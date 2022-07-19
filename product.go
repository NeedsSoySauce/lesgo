package main

type Product struct {
	Name         string `json:"name"`
	PriceInCents uint64 `json:"priceInCents"`
}

type ProductEntity struct {
	ID uint `json:"id"`
	Product
}

func NewProduct(id uint, name string, priceInCents uint64) ProductEntity {
	product := Product{name, priceInCents}
	return ProductEntity{id, product}
}
