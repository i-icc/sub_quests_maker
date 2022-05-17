package controller

import (
  "net/http"
  "os"
  "fmt"
	
	"github.com/labstack/echo/v4"
  "github.com/bradrydzewski/go.auth"
  "github.com/joho/godotenv"
)

func login() echo.HandlerFunc {
  return func(c echo.Context) error {
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
    auth.Config.LoginSuccessRedirect = "/api/test"
    auth.Config.CookieSecure = false

    twitterCallBack := "/auth/login"
    twitterHandler := auth.Twitter(API_KEY, API_KEY_SECRET, twitterCallBack)

    return c.Handler(http.StatusOK, twitterHandler)
  }
}
