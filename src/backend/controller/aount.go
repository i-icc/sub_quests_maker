package controller

import (
  "net/http"
  "encoding/json"
  "fmt"
	
  "github.com/bradrydzewski/go.auth"
)

func Login(w http.ResponseWriter, r *http.Request) {
  token := r.FormValue("oauth_token")
  if token != ""{
    // fmt.Println("oauth_token:", token)
    data := map[string]interface{}{
      "token": token,
    }
    j, err := json.Marshal(data)
    if err != nil {
      fmt.Println(err)
      return
    }
    fmt.Fprintf(w, string(j))
    return
  } else {
    http.Redirect(w, r, "/auth/accese", http.StatusSeeOther)
  }
}

// logout handler
func Logout(w http.ResponseWriter, r *http.Request) {
	auth.DeleteUserCookie(w, r)
	http.Redirect(w, r, "/api/usertest", http.StatusSeeOther)
}
