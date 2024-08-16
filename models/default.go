package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DefaultModel struct {
    ID   		string 		`gorm:"primary_key;" json:"id"`
	CreatedAt 	time.Time 	`json:"createdAt"`
	UpdatedAt 	time.Time 	`json:"updatedAt"`
}

func (b *DefaultModel) BeforeCreate(tx *gorm.DB) (err error) {
	uuid,err := uuid.NewV7();
	b.ID = uuid.String();
	return 
}