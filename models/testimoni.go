// models/teacher.go
package models

type Testimoni struct {
	DefaultModel
	Name  		string			`json:"name"`
	Testimoni	string			`json:"testimoni"`
	Image 		string			`json:"image"`
	Item    	Item   			`json:"item"`
	ItemID  	string 			`json:"itemID"`
	User    	User      		`json:"user"`
	UserID  	string    		`json:"userID"`
}
