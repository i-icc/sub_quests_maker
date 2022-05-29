package controller

import (
	_ "bytes"
	"encoding/json"
	"fmt"
	_ "log"
	"net/http"

	"backend/db"
)

func CreateQuest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var params map[string]bool

		params["when"] = r.FormValue("when") == "on"
		params["where"] = r.FormValue("where") == "on"
		params["who"] = r.FormValue("who") == "on"
		params["what"] = r.FormValue("what") == "on"

		if !params["what"] {
			fmt.Fprintf(w, "")
			return
		}

		db := db.Connect()
		defer db.Close()

		var user []interface{}

		j, err := json.Marshal(user)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Fprintf(w, string(j))
	default:
		fmt.Fprintf(w, "Method not allowd")
	}
}

func GetRandomInstruction(table string) interface{} {
	type Result struct {
		Id          string
		Instruction string
	}
	result := Result{}

	return result
}

// func ResistQuest(u model.Quest) {
// 	db := db.Connect()
// 	defer db.Close()

// 	db.Create(&u)

// 	j, err := json.Marshal(u)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	log.Print(string(j))
// }
