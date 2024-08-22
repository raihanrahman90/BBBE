package models

import (
	"bbbe/enums"
)

type History struct {
	DefaultModel
	DataId 		int				`json:"dataId"`
	DataType 	enums.DataType	`json:"dataType"`
}
