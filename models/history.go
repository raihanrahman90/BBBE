package models

import (
	"rumahbelajar/enums"
)

type History struct {
	DefaultModel
	DataId 		int				`json:"dataId"`
	DataType 	enums.DataType	`json:"dataType"`
}
