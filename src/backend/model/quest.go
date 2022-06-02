package model

import (
	"github.com/jinzhu/gorm"
)

type Quest struct {
	gorm.Model
	Id       string `json:"id"`
	User_uid string `json:"user_uid"`
	Who_id   string `json:"who_id"`
	Where_id string `json:"where_id"`
	What_id  string `json:"what_id"`
	When_id  string `json:"when_id"`
	Comment  string `json:"comment"`
}
