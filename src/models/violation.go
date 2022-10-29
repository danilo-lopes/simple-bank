package models

type Violation struct {
	Account    Account  `json:"account,required"`
	Violations []string `json:"violations"`
}
