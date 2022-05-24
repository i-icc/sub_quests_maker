package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	_ "os"
	_ "strings"

	"backend/model"
)

func GetAcount(token string) model.User {
	url := "https://api.twitter.com/2/users/me?user.fields=profile_image_url"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	client := new(http.Client)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		panic(err.Error())
	}
	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))

	var data map[string]model.User
	if err := json.Unmarshal([]byte(string(byteArray)), &data); err != nil {
		fmt.Println(err)
	}

	return data["data"]
}
