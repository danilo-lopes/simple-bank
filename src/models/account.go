package models

import (
	"encoding/json"
	"fmt"
)

type Account struct {
	ID             uint64        `json:"id,omitempty"`
	Active         bool          `json:"active,required"`
	AvailableLimit int           `json:"availableLimit,required"`
	History        []Transaction `json:"transactions,omitempty"`
}

func (a Account) Details() error {
	model, erro := json.Marshal(a)
	if erro != nil {
		return erro
	}

	fmt.Printf("%s\n", string(model))
	return nil
}
