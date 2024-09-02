// models/teacher.go
package models

type Item struct {
	DefaultModel
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Category	*string `json:"category"`
	Description string `json:"description"`
	Image       string `json:"image"`
}
