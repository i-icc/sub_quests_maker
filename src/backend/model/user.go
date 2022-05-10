package model

import (
    "github.com/jinzhu/gorm"
)

type User struct {
    gorm.Model
    //`json:"--"`はリクエストボディに指定する名前
	Id                  int    `json:"id"`
    Uid                 string `json:"uid"`
    Nickname            string `json:"nickname"`
}