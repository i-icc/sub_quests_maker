package controller

import (
	"fmt"
	"log"
	"net/http"

	auth "github.com/bradrydzewski/go.auth"
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

func SetCookie(w http.ResponseWriter, r *http.Request) {
	// cookie := http.Cookie{
	// 	Name:   "tokenId",
	// 	Value:  "abc",
	// 	MaxAge: 60 * 60,
	// }
	// http.SetCookie(w, &cookie)

	fmt.Fprintf(w, "Cookieの設定ができたよ")

	// http.Redirect(w, r, "/auth/login", 302)
}

func DeleteCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("tokenId")

	if err != nil {
		log.Println("Cookie: ", err)
		return
	}
	cookie.MaxAge = -1
	// http.SetCookie(w, cookie)

	// fmt.Fprintf(w, "消したで\n")
	http.Redirect(w, r, "/api/cookie/set", http.StatusSeeOther)
}

func ShowCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("tokenId")

	if err != nil {
		log.Println("Cookie: ", err)
	}

	fmt.Fprintf(w, cookie.Value)
}
