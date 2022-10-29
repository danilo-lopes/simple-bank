package engine

import "github.com/danilo-lopes/simple-bank/src/models"

func Debit(t models.Transaction, a *models.Account) {
	value := a.AvailableLimit - t.Amount
	a.AvailableLimit = value
}
