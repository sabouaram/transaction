package models

type User struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

type Transaction struct {
	FromId string  `json:"from_id"`
	ToId   string  `json:"to_id"`
	Amount float64 `json:"amount"`
}
