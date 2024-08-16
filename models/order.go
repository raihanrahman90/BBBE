// models/teacher.go
package models

import "time"

type Order struct {
	DefaultModel
	User    User      `json:"user"`
	UserID  string    `json:"userID"`
	Payment string    `json:"payment"`
	Date    time.Time `json:"date"`
	Status  string    `json:"status"`
}
