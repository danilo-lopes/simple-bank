package engine

import (
	"fmt"

	"github.com/danilo-lopes/simple-bank/src/models"
)

func AuthorizeTransaction(t models.Transaction, a *models.Account) models.Violation {
	var violation models.Violation

	if err := validateAccount(*a); err != nil {
		violation.Violations = append(violation.Violations, err.Error())
	}

	if err := validateLimit(t, *a); err != nil {
		violation.Violations = append(violation.Violations, err.Error())
	}

	if len(violation.Violations) > 0 {
		violation.Account = *a
		return violation
	}

	violation.Account = *a
	Debit(t, a)

	violation.Account.History = append(violation.Account.History, t)
	violation.Violations = append(violation.Violations, violation.Violations...)

	return violation
}

func AuthorizeAccountCreation(u models.User) models.Violation {
	var violation models.Violation

	if err := u.UserHaveName(); err != nil {
		violation.Violations = append(violation.Violations, err.Error())
	}

	if len(violation.Violations) > 0 {
		violation.Account = u.Account
		return violation
	}

	return violation
}

func validateAccount(a models.Account) error {
	if !a.Active {
		return fmt.Errorf("account is not active")
	}

	return nil
}

func validateLimit(t models.Transaction, a models.Account) error {
	if t.Amount > a.AvailableLimit {
		return fmt.Errorf("the amount is higher then the client limit")
	}

	return nil
}
