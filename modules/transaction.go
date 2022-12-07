package modules

type Transaction struct {
	Items []TransactionItem
	Money int
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

func (transaction Transaction) AddItem(item TransactionItem) {
	transaction.Items = append(transaction.Items, item)
}
