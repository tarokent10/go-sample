package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("aa")
	fmt.Fprint(w, "Hello, HTTPサーバ")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
	fmt.Println("aa")
}
