package controller

import (
  "net/http"
  "encoding/json"
  "fmt"
	
  "github.com/bradrydzewski/go.auth"
)

func Siginup(w http.ResponseWriter, r *http.Request) {
  token := r.FormValue("oauth_token")
  if token != ""{
  }
}

func Login(w http.ResponseWriter, r *http.Request) {
  oauth_token := r.FormValue("oauth_token")
  oauth_verifier := r.FormValue("oauth_verifier")
  if oauth_token != "" && oauth_verifier != "" {
    // fmt.Println(GetAccessToken(oauth_token, oauth_verifier))

    data := map[string]interface{}{
      "oauth_token": oauth_token,
      "oauth_verifier": oauth_verifier,
    }
    j, err := json.Marshal(data)
    if err != nil {
      fmt.Println(err)
      return
    }
    fmt.Fprintf(w, string(j))
    
    // GetAcount(token)
  } else {
    http.Redirect(w, r, "/auth/access", http.StatusSeeOther)
  }
}

// logout handler
func Logout(w http.ResponseWriter, r *http.Request) {
	auth.DeleteUserCookie(w, r)
	http.Redirect(w, r, "/api/usertest", http.StatusSeeOther)
}
