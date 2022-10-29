package main

// func TestAccountActive(t *testing.T) {
// 	joao := Account{
// 		Active:         true,
// 		AvailableLimit: 10,
// 		History:        nil,
// 	}

// 	if err := accountValidator(*&joao); err != "" {
// 		t.Errorf(err)
// 	}
// }

// func TestAccountInactive(t *testing.T) {
// 	joao := Account{
// 		Active:         false,
// 		AvailableLimit: 10,
// 		History:        nil,
// 	}

// 	if err := accountValidator(*&joao); err == "" {
// 		t.Errorf(err)
// 	}
// }

// func TestAccountTransaction(t *testing.T) {
// 	joao := Account{
// 		Active:         true,
// 		AvailableLimit: 10,
// 		History:        nil,
// 	}
// 	joaoTransaction := Transaction{
// 		Amount:   2,
// 		Merchant: "Mercado livre",
// 		Time:     time.Now(),
// 	}
// }

// func TestAccountTransactionLimit(t *testing.T) {
// 	joao := Account{
// 		Active:         true,
// 		AvailableLimit: 10,
// 		History:        nil,
// 	}
// 	joaoTransaction := Transaction{
// 		Amount:   11,
// 		Merchant: "Mercado livre",
// 		Time:     time.Now(),
// 	}
// }
