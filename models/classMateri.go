// models/teacher.go
package models

type ClassModules struct {
	DefaultModel
	Module 		string		`json:"module"`
	ClassID 	string		`json:"classId"`
	Class 		Class 		`gorm:"foreignkey:ClassID"`
}
