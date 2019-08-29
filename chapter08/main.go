package main

import (
	"net/http"

	"github.com/nsiregar/solid-go/chapter08/di"
)

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	di.Greet(w, "world")
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(GreetHandler))
}
