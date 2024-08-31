// models/teacher.go
package models

type Item struct {
	DefaultModel
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Image       string `json:"image"`
}
