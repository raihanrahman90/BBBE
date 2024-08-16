// models/teacher.go
package models

type Class struct {
	DefaultModel
	Title 		string			`json:"title"`
	Price 		int				`json:"price"`
	Description string			`json:"description"`
	Image		string			`json:"image"`
	Modules 	[]ClassModules	`json:"modules"`
}
