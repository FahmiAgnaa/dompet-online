package resp

import (
	"time"
)

type GetTransactionsResponse struct {
	Id            string    `json:"id_transaction"`
	Destination   string    `json:"destination"`
	Amount        int       `json:"amount"`
	Description   string    `json:"description"`
	CreateAt      time.Time `json:"time_of_transaction"`
	User          User      `json:"user"`
	Wallet        Wallet    `json:"wallet"`
	PaymentMethod PaymentMethod
}

type PaymentMethod struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type User struct {
	UserName string `json:"user_name"`
}

type Wallet struct {
	RekeningUser string `json:"rekening_user"`
	Balance      int    `json:"balance"`
}
