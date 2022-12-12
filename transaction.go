package main

import "fmt"

type Transaction struct {
	Items []TransactionItem
	Money int
}

func NewTransaction() *Transaction {
	return &Transaction{Items: []TransactionItem{}, Money: 0}
}

func (transaction Transaction) GetTotal() int {
	total := 0

	for _, item := range transaction.Items {
		total += item.GetSubTotal()
	}

	return total
}

func (transaction Transaction) GetChange() int {
	return transaction.Money - transaction.GetTotal()
}

func (transaction *Transaction) AddItem(item TransactionItem) {
	transaction.Items = append(transaction.Items, item)
}

func (transaction Transaction) PrintStruct() {
	for index, item := range transaction.Items {
		fmt.Println(index+1, ". ", item.Product.Name, " : Rp.", item.Product.Price, " x ", item.Amount, " = ", item.GetSubTotal())
	}
	fmt.Println("Total : Rp. ", transaction.GetTotal())
}
