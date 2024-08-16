// models/teacher.go
package models

type OrderItem struct {
	DefaultModel
	Order   Order  `json:"order"`
	OrderID string `json:"orderID"`
	Price   int    `json:"price"`
	Amount  int    `json:"amount"`
}
