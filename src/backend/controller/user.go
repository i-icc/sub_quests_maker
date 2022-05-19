package controller

import (
	"net/http"
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"backend/db"
	"backend/model"
)

func GetUserTest(w http.ResponseWriter, r *http.Request){
	db := db.Connect()
	defer db.Close()

	var user []model.User
    db.Raw("SELECT * FROM user").Scan(&user)

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(&user); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())

	_, err := fmt.Fprint(w, buf.String())
	if err != nil {
		return
	}
}