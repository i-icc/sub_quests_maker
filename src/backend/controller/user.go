package controller

import (
	_ "bytes"
	"encoding/json"
	"fmt"
	"log"
	_ "log"
	"net/http"

	"backend/db"
	"backend/model"
)

func GetUserTest(w http.ResponseWriter, r *http.Request) {
	db := db.Connect()
	defer db.Close()

	var user []model.User
	db.Raw("SELECT * FROM user").Scan(&user)

	j, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, string(j))
}

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
