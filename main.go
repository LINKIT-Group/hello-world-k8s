package main

import "io"
import "net/http"

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World - v2")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}
