package controller

import (
	"net/http"
  "fmt"

  "github.com/bradrydzewski/go.auth"
)

func GetTest(w http.ResponseWriter, r *http.Request, u auth.User) {
  fmt.Fprintf(w, "test\n") 

	fmt.Fprintf(w, u.Id()+"\n")
  fmt.Fprintf(w, u.Provider()+"\n")
  fmt.Fprintf(w, u.Name()+"\n")
  fmt.Fprintf(w, u.Email()+"\n")
  fmt.Fprintf(w, u.Org()+"\n")
  fmt.Fprintf(w, u.Link()+"\n")
  fmt.Fprintf(w, u.Picture()+"\n")
}