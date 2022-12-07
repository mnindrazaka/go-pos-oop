package main

type TransactionItem struct {
	Product Product
	Amount  int
}

func NewTransactionItem(product Product, amount int) TransactionItem {
	return TransactionItem{product, amount}
}

func (item TransactionItem) GetSubTotal() int {
	return item.Product.Price * item.Amount
}
