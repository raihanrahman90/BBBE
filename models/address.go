// models/teacher.go
package models

type Address struct {
	DefaultModel
	User    	User      `json:"user"`
	UserID  	string    `json:"userID"`
	Address 	string    `json:"address"`
	City		string	  `json:"city"`
	Province	string	  `json:"province"`
}