package dbl

import "github.com/bojanz/currency"

func addMoney(amount1, amount2 currency.Amount) (currency.Amount, error) {
	return amount1.Add(amount2)
}
