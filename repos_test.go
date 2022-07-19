package main

import (
	"reflect"
	"testing"
)

func TestInMemoryProductsRepository_GetProducts(t *testing.T) {
	tests := []struct {
		name    string
		repo    *InMemoryProductsRepository
		want    []ProductEntity
		wantErr bool
	}{
		{"empty", &InMemoryProductsRepository{products: []ProductEntity{}}, []ProductEntity{}, false},
		{"should fail", &InMemoryProductsRepository{products: []ProductEntity{}}, []ProductEntity{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.repo.GetProducts()
			if (err != nil) != tt.wantErr {
				t.Errorf("InMemoryProductsRepository.GetProducts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InMemoryProductsRepository.GetProducts() = %v, want %v", got, tt.want)
			}
		})
	}
}
