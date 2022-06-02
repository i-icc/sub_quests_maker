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

func CreateQuest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		params := map[string]bool{}
		param2table := map[string]string{
			"when":  "timings",
			"where": "places",
			"who":   "whos",
			"what":  "whats",
		}

		params["when"] = r.FormValue("when") == "on"
		params["where"] = r.FormValue("where") == "on"
		params["who"] = r.FormValue("who") == "on"
		params["what"] = r.FormValue("what") == "on"

		if !params["what"] {
			fmt.Fprintf(w, "必須パラメーが無いで")
			return
		}

		var quest []interface{}

		for k, v := range param2table {
			if params[k] {
				quest = append(quest, GetRandomInstruction(v))
			}
		}

		j, err := json.Marshal(quest)
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
		Tag         string `json:"tag"`
		Id          string `json:"id"`
		Instruction string `json:"instruction"`
	}

	db := db.Connect()
	defer db.Close()

	var result Result
	db.Raw("SELECT * FROM " + table + " ORDER BY RAND() LIMIT 1;").Scan(&result)
	result.Tag = table

	return result
}

func ResistQuest(u model.Quest) {
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
