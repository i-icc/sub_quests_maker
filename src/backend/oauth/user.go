package oauth

import (
	"encoding/json"
	"fmt"
	"log"

	"backend/db"
	"backend/model"
)

func CheckExsistUser(u model.User) bool {
	db := db.Connect()
	defer db.Close()

	var user []model.User
	db.Raw("SELECT * FROM users WHERE uid = ? LIMIT 1", u.Uid).Scan(&user)

	var result bool
	if len(user) == 1 {
		result = true
	} else {
		result = false
	}
	return result

}

func ResistUser(u model.User) {
	db := db.Connect()
	defer db.Close()

	db.Create(&u)

	j, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Print(string(j))
}
