package controller

import (
	"net/http"
  "fmt"
)

func GetTest(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "test") //ここでwに入るものがクライアントに出力されます。
}