// models/teacher.go
package models

type TeacherField struct {
	DefaultModel
	NameField 	string		`json:"nameField"`
	TeacherID 	string		`json:"teacherId"`
	Teacher 	Teacher 	`gorm:"foreignkey:TeacherID"`
}
