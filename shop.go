package main

import "fmt"

type Shop struct {
	Products    []Product
	Transaction []Transaction
}

func (shop Shop) AddProduct(product Product) {
	shop.Products = append(shop.Products, product)
}

func (shop Shop) PrintProductList() {
	fmt.Println("Products list : ")
	for index, product := range shop.Products {
		fmt.Println((index + 1), ". ", product.Name, " : Rp.", product.Price)
	}
}

func (shop Shop) AddTransaction(product Product) {
	shop.Products = append(shop.Products, product)
}

func (shop Shop) PrintLogTransactions(product Product) {
	shop.Products = append(shop.Products, product)
}
