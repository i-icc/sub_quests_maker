package main

import (
	"net/http"
	"log"

	"backend/route"
)

func main() {
	e := route.Init()
	log.Fatal(http.ListenAndServe(":3000", e))
}
