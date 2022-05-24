package model

import (
	"github.com/jinzhu/gorm"
)

type What struct {
	gorm.Model
	//`json:"--"`はリクエストボディに指定する名前
	Id          string `json:"id" xml:"id"`
	Instruction string `json:"instruction" xml:"instruction"`
}
