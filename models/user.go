// models/user.go
package models

import (
	"bbbe/enums"
)

type User struct {
	DefaultModel
	Username 		string 			`json:"username";gorm:"unique"`
	Password 		string			`json:"password"`
	Access 			enums.Access	`json:"access"`
	RefreshToken 	string			`json:"refreshToken"`
}
