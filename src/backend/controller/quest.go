package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"backend/db"
	"backend/model"
	"backend/oauth"
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
	case "POST":
		cookie, err := r.Cookie("tokenId")
		if err != nil {
			fmt.Println("Cookie: ", err)
			fmt.Fprintf(w, "ログインし直してください")
			return
		}
		fmt.Print(cookie.Value)
		tokenId := cookie.Value
		if !oauth.Mg.Exists(tokenId) {
			fmt.Fprintf(w, "ログインし直してください")
			return
		}
		var u model.Quest
		u.User_uid = oauth.Mg.GetToken(tokenId).GetUid()
		u.When_id, _ = strconv.Atoi(r.FormValue("timings"))
		u.Where_id, _ = strconv.Atoi(r.FormValue("places"))
		u.Who_id, _ = strconv.Atoi(r.FormValue("whos"))
		u.What_id, _ = strconv.Atoi(r.FormValue("whats"))

		resistQuest(u)
	default:
		fmt.Fprintf(w, "Method not allowd")
	}
}

func GetRandomInstruction(table string) interface{} {
	type Result struct {
		Tag         string `json:"tag"`
		Id          int    `json:"id"`
		Instruction string `json:"instruction"`
	}

	db := db.Connect()
	defer db.Close()

	var result Result
	db.Raw("SELECT * FROM " + table + " ORDER BY RAND() LIMIT 1;").Scan(&result)
	result.Tag = table

	return result
}

func resistQuest(u model.Quest) {
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
