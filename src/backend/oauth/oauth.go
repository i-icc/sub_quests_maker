package oauth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/joho/godotenv"

	"backend/controller"
	_ "backend/model"
)

const (
	response_type = "code"
	redirect_uri  = "http://localhost:3000/auth/callback"
	grant_type    = "authorization_code"

	verifier = "E9Melhoa2OwvFrEMTJguCHaoeK1t8URWbuGJSstw-cM"
)

type Oauth struct {
	clientId              string
	clientSecret          string
	scope                 string
	state                 string
	code_challenge_method string
	code_challenge        string
	authEndpoint          string
	tokenEndpoint         string
}

func (oauth *Oauth) SetUp() {
	err := godotenv.Load("/backend/.env")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("env読み取り成功")
	}

	oauth.clientId = os.Getenv("CLIENT_ID")
	oauth.clientSecret = os.Getenv("CLIENT_SECRET")
	oauth.authEndpoint = "https://twitter.com/i/oauth2/authorize?"
	oauth.tokenEndpoint = "https://api.twitter.com/2/oauth2/token"
	oauth.state = createRandomState()
	oauth.scope = "tweet.read tweet.write users.read offline.access"
	oauth.code_challenge_method = "S256"
	oauth.code_challenge = base64URLEncode()
}

func base64URLEncode() string {
	hash := sha256.Sum256([]byte(verifier))
	return base64.RawURLEncoding.EncodeToString(hash[:])
}

func (oauth *Oauth) Login(w http.ResponseWriter, r *http.Request) {

	authEndpoint := oauth.authEndpoint

	values := url.Values{}
	values.Add("scope", oauth.scope)
	values.Add("redirect_uri", redirect_uri)
	values.Add("response_type", response_type)
	values.Add("client_id", oauth.clientId)
	values.Add("state", oauth.state)
	values.Add("code_challenge_method", oauth.code_challenge_method)
	values.Add("code_challenge", oauth.code_challenge)

	http.Redirect(w, r, authEndpoint+values.Encode(), 302)
}

func (oauth *Oauth) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("tokenId")
	if err != nil {
		fmt.Println("Cookie: ", err)
		http.Redirect(w, r, "/auth/login", http.StatusBadRequest)
		return
	}
	tokenId := cookie.Value

	c := http.Cookie{
		Name:   "tokenId",
		MaxAge: -1,
		Path:   "/",
	}
	http.SetCookie(w, &c)

	if !mg.Exists(tokenId) {
		fmt.Println("Not Found Session")
		http.Redirect(w, r, "/auth/login", http.StatusBadRequest)
	}
	mg.Destroy(tokenId)

	http.Redirect(w, r, "/api/usertest", http.StatusSeeOther)
}

// func (oauth *Oauth) Siginup(w http.ResponseWriter, r *http.Request) {
//
// }

func (oauth *Oauth) Callback(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	if query.Get("state") != oauth.state {
		w.WriteHeader(403)
		w.Write([]byte("status error"))
		return
	}

	result, err := oauth.tokenRequest(query)
	if err != nil {
		log.Println(err)
	}

	// fmt.Fprintf(w, result["access_token"])

	u := controller.GetAcount(result["access_token"])
	if !controller.CheckExsistUser(u) {
		controller.ResistUser(u)
	}

	t := NewToken(result["access_token"], u.Uid)
	log.Println(t.id, t.uid, t.token)
	mg.Save(t)

	c := http.Cookie{
		Name:   "tokenId",
		Value:  t.id,
		MaxAge: 60 * 60 * 24,
		Path:   "/",
	}
	http.SetCookie(w, &c)

	fmt.Fprintf(w, t.id)
	// http.Redirect(w, r, "/api/cookie/check", http.StatusSeeOther)
}

func (oauth *Oauth) tokenRequest(query url.Values) (map[string]string, error) {

	tokenEndpoint := oauth.tokenEndpoint
	values := url.Values{}
	values.Add("code", query.Get("code"))
	values.Add("grant_type", grant_type)
	values.Add("client_id", oauth.clientId)
	values.Add("redirect_uri", redirect_uri)
	values.Add("code_verifier", verifier)

	r, err := http.NewRequest("POST", tokenEndpoint, strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}
	r.SetBasicAuth(oauth.clientId, oauth.clientSecret)
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		log.Printf("request err: %s", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// log.Printf("token response : %s", string(body))
	// var data map[string]interface{}
	var data map[string]string
	json.Unmarshal(body, &data)

	return data, nil
}

func createRandomState() string {
	b := make([]byte, 64)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
