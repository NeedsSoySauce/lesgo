package main

import (
	"encoding/json"
	"net/http"
)

type ProductsApplication struct {
	repository ProductsRepository
}

func NewProductsApplication(repository ProductsRepository) ProductsApplication {
	return ProductsApplication{repository}
}

func (app ProductsApplication) HandleRequest(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		app.getProducts(w, req)
	case http.MethodPost:
		app.createProduct(w, req)
	}
}

func (app ProductsApplication) getProducts(w http.ResponseWriter, req *http.Request) {
	products, err := app.repository.GetProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	response, err := json.Marshal(products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
}

func (app ProductsApplication) createProduct(w http.ResponseWriter, req *http.Request) {
	var product Product

	err := json.NewDecoder(req.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	productEntity, err := app.repository.AddProduct(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	response, err := json.Marshal(productEntity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
}
