package oauth

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

type Token struct {
	id    string
	token string
	uid   string
}

func NewToken(t string, u string) *Token {
	return &Token{
		id:    createId(),
		token: t,
		uid:   u,
	}
}

func createId() string {
	b := make([]byte, 64)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
