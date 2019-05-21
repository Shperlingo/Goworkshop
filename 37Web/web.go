package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world from web")
}
func index1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello about world from web")
}
func main() {
	fmt.Println("Hello World")
	http.HandleFunc("/", index)
	http.HandleFunc("/about", index1)
	http.ListenAndServe(":3000", nil)
}
