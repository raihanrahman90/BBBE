// models/teacher.go
package models

type LandingPage struct {
	DefaultModel
	Key 	string	`json:"key"`
	Type    string	`json:"type"`	
	Value 	string	`json:"value"`
}
