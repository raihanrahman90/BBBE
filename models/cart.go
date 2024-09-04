// models/teacher.go
package models

type Cart struct {
	DefaultModel
	Item   	Item  		`json:"item"`
	ItemID 	string 		`json:"itemID"`
	User    User      	`json:"user"`
	UserID  string    	`json:"userID"`
}
