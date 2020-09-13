package main

import (
	"fmt"
	"net/http"
	"strings"
)

func hello(w http.ResponseWriter, r *http.Request) {
	var segments = strings.Split(r.URL.Path, "/")

	fmt.Fprintf(w, "Hello %v!", segments[len(segments) - 1])
}

func main() {
	http.HandleFunc("/hello/", hello)

	println("Listening on port 3333")
	http.ListenAndServe(":3333", nil)
}
