// models/teacher.go
package models

type ItemImage struct {
	DefaultModel
	Path   string `json:"path"`
	ItemID string `json:"itemId"`
	Item   Item   `gorm:"foreignkey:ItemID"`
}
