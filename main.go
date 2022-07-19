package main

import (
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	products := []ProductEntity{
		NewProduct(1, "Chrome Toaster", 10000),
		NewProduct(2, "Copper Kettle", 4999),
		NewProduct(3, "Mixing Bowl", 2000),
	}
	repository, err := NewRemoteProductsRepository(products, nil)

	if err != nil {
		log.Fatal(err)
	}

	// log.Println(repository.GetProducts())
	// log.Println(repository.AddProduct(Product{"asdf", 124}))
	// log.Println(repository.GetProducts())

	application := NewProductsApplication(repository)

	server := http.NewServeMux()

	server.HandleFunc("/products", application.HandleRequest)

	http.ListenAndServe(":1234", server)
}
