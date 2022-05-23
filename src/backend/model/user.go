package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	//`json:"--"`はリクエストボディに指定する名前
	Uid      string `json:"username" xml:"username"`
	Nickname string `json:"name" xml:"name"`
	Image    string `json:"profile_image_url" xml:"profile_image_url"`
}
