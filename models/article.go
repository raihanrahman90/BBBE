// models/teacher.go
package models

type Article struct {
	DefaultModel
	Title 		string		`json:"title"`
	Body 		string		`json:"description"`
	Image		string		`json:"image"`
}
