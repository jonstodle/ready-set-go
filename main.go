package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func hello(w http.ResponseWriter, r *http.Request) {
	var segments = strings.Split(r.URL.Path, "/")[2:]

	json.NewEncoder(w).Encode(struct {
		Greeting string `json:"greeting"`
		Names []string `json:"names"`
	}{
		"Hello", segments,
	})
}

func main() {
	http.HandleFunc("/hello/", hello)

	println("Listening on port 3333")
	http.ListenAndServe(":3333", nil)
}
