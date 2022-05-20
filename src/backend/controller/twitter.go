package controller

import (
  "os"
  "fmt"
  "net/http"
  _ "encoding/json"
  "io/ioutil"
  _ "strings"
  
  _ "backend/model"
	
  "github.com/bradrydzewski/go.auth"
  "github.com/joho/godotenv"
)

func Access() *auth.AuthHandler {
  // keyを .envファイルから読み取る
  err := godotenv.Load("/backend/.env")
  if err != nil {
      panic(err.Error())
  } else {
      fmt.Println("env読み取り成功")
  }

  API_KEY := os.Getenv("API_KEY")
  API_KEY_SECRET := os.Getenv("API_KEY_SECRET")

  // set the auth parameters
  auth.Config.CookieSecret = []byte("7H9xiimk2QdTdYI7rDddfJeV")
  auth.Config.LoginSuccessRedirect = "/auth/login"
  auth.Config.CookieSecure = false

  twitterCallBack := "http://localhost:3000/auth/login"
  twitterHandler := auth.Twitter(API_KEY, API_KEY_SECRET, twitterCallBack)

  return twitterHandler
}

func GetAcount(token string) { //model.User {
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
	// return model.User("","","","")
	// https://twitter.com/i/oauth2/authorize?response_type=code&client_id=<Client ID>&redirect_uri=https://127.0.0.1:3000/cb&scope=tweet.read%20users.read%20offline.access&state=abc&code_challenge=E9Melhoa2OwvFrEMTJguCHaoeK1t8URWbuGJSstw-cM&code_challenge_method=s256
}

func GetAccessToken(oauth_token string, oauth_verifier string) string {
	url := "https://api.twitter.com/oauth/access_token?"
	url += "oauth_token="+oauth_token
	url += "&oauth_verifier="+ oauth_verifier
	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Set("content-type", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		panic(err.Error())
	}
	byteArray, _ := ioutil.ReadAll(resp.Body)
  	return string(byteArray)
}