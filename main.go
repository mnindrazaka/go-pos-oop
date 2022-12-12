package main

import (
	"fmt"
	"strings"
)

func main() {
	shop := NewShop([]Product{
		{Name: "Pencil", Price: 3000},
		{Name: "Book", Price: 3500},
		{Name: "Eraser", Price: 2000},
	}, []Transaction{})

	for {
		fmt.Println("========================================")
		fmt.Println("Welcome to Our Shop")
		fmt.Println("Options : ")
		fmt.Println("1. Add Product")
		fmt.Println("2. Create Transaction")
		fmt.Println("3. Exit")
		fmt.Println("========================================")

		fmt.Print("Enter your option : ")
		var option int
		fmt.Scanf("%d", &option)

		if option == 1 {
			fmt.Print("Enter product name : ")
			var name string
			fmt.Scanf("%s", &name)

			fmt.Print("Enter product price : ")
			var price int
			fmt.Scanf("%d", &price)

			shop = shop.AddProduct(Product{Name: name, Price: price})

		} else if option == 2 {
			transaction := NewTransaction()

			for {
				shop.PrintProductList()

				fmt.Print("Enter product option : ")
				var option int
				fmt.Scanf("%d", &option)

				product := shop.Products[option-1]

				fmt.Print("Enter product amount : ")
				var amount int
				fmt.Scanf("%d", &amount)

				transaction = transaction.AddItem(NewTransactionItem(product, amount))

				fmt.Print("Do you want to add another product (y/n) ? ")
				var repeatAnswer string
				fmt.Scanf("%s", &repeatAnswer)

				if strings.ToLower(repeatAnswer) != "y" {
					break
				}
			}

			transaction.PrintStruct()

			fmt.Print("Enter your money : ")
			fmt.Scanf("%d", &transaction.Money)

			if transaction.Money >= transaction.GetTotal() {
				fmt.Println("your change is : Rp.", transaction.GetChange())
			} else {
				fmt.Println("your money is not enough")
			}
		} else if option == 3 {
			fmt.Println("Bye bye")
			break
		} else {
			fmt.Println("wrong option")
		}
	}
}
