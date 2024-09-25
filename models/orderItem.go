// models/teacher.go
package models

type OrderItem struct {
	DefaultModel
	Order   Order  `json:"order"`
	OrderID string `json:"orderID"`
	Name    string `json:"name"`
	Price   int    `json:"price"`
	Amount  int    `json:"amount"`
	Image   string `json:"image"`
}
