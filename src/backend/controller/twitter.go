package controller

import (
  "os"
  "fmt"
	
  "github.com/bradrydzewski/go.auth"
  "github.com/joho/godotenv"
)

func Accese() *auth.AuthHandler {
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