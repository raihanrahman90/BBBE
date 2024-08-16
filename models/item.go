// models/teacher.go
package models

type Item struct {
	DefaultModel
	Name        string      `json:"name"`
	Price       int         `json:"price"`
	Description string      `json:"description"`
	ItemImage   []ItemImage `json:"item_image" gorm:"constraints:OnDelete:Cascade;"`
}
