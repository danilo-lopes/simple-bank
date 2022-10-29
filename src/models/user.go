package models

import (
	"errors"
)

type User struct {
	ID      uint64  `json:"id,omitempty"`
	Name    string  `json:"name,required"`
	Email   string  `json:"email,required"`
	CPF     string  `json:"cpf,required"`
	Account Account `json:"account,omitempty"`
}

func (u User) UserHaveName() error {
	if u.Name == "" {
		return errors.New("the user doesnt have name")
	}

	return nil
}
