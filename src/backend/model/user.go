package model

import (
    "github.com/jinzhu/gorm"
)

type User struct {
    gorm.Model
    //`json:"--"`はリクエストボディに指定する名前
	Id int `json:"id" xml:"id"`
    Uid string `json:"uid" xml:"uid"`
    Nickname string `json:"nickname" xml:"nickname"`
}