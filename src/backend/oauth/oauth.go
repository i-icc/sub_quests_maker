package oauth

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"os"

	"github.com/joho/godotenv"

	_ "backend/controller"
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
	
	oauth.clientId =  os.Getenv("CLIENT_ID")
	oauth.clientSecret = os.Getenv("CLIENT_SECRET")
	oauth.authEndpoint = "https://twitter.com/i/oauth2/authorize?"
	oauth.tokenEndpoint = "https://api.twitter.com/2/oauth2/token"
	oauth.state = "abc" // ここはランダムに
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

func (oauth *Oauth) Callback(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	if query.Get("state") != oauth.state {
		w.WriteHeader(403)
		w.Write([]byte("status error"))
	}

	result, err := oauth.tokenRequest(query)
	if err != nil {
		log.Println(err)
	}
	j, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintf(w, string(j))
}

// 認可コードを使ってトークンリクエストをエンドポイントに送る
func (oauth *Oauth) tokenRequest(query url.Values) (map[string]interface{}, error) {

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

	log.Printf("token response : %s", string(body))
	var data map[string]interface{}
	json.Unmarshal(body, &data)

	return data, nil
}