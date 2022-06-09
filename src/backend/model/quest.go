package model

import (
	"github.com/jinzhu/gorm"
)

type Quest struct {
	gorm.Model
	Id       int    `json:"id"`
	User_uid string `json:"user_uid"`
	Who_id   int    `json:"who_id"`
	Where_id int    `json:"where_id"`
	What_id  int    `json:"what_id"`
	When_id  int    `json:"when_id"`
	Comment  string `json:"comment"`
}
