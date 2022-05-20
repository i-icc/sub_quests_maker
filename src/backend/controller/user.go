package controller

import (
	"net/http"
	_ "bytes"
	"encoding/json"
	"fmt"
	_ "log"

	"backend/db"
	"backend/model"
)

func GetUserTest(w http.ResponseWriter, r *http.Request){
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