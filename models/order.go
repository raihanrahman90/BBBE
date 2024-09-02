// models/teacher.go
package models

import "time"

type Order struct {
	DefaultModel
	User    	User      `json:"user"`
	UserID  	string    `json:"userID"`
	OrderItem	[]OrderItem	`json:"order_item"`
	Payment 	string    `json:"payment"`
	Date    	time.Time `json:"date"`
	Status  	string    `json:"status"`
	ProofOfPayment	string	`json:"proof_of_payment"`
	Address		string	  `json:"address"`
	City		string	  `json:"city"`
	Province	string		`json:"province"`
}
