package main

import "fmt"

type Shop struct {
	Products    []Product
	Transaction []Transaction
}

func NewShop(products []Product, transactions []Transaction) *Shop {
	return &Shop{Products: products, Transaction: transactions}
}

func (shop *Shop) AddProduct(product Product) {
	shop.Products = append(shop.Products, product)
}

func (shop *Shop) PrintProductList() {
	fmt.Println("Products list : ")
	for index, product := range shop.Products {
		fmt.Println((index + 1), ". ", product.Name, " : Rp.", product.Price)
	}
}
