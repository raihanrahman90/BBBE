// models/teacher.go
package models

type Teacher struct {
	DefaultModel
	Name  		string			`json:"name"`
	Description	string			`json:"description"`
	Image 		string			`json:"image"`
	Fields 		[]TeacherField	`json:"fields"`
}
