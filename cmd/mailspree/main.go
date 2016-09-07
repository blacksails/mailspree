package main

import "github.com/blacksails/mailspree/http"

func main() {
	s := http.Server{}
	http.ListenAndServe(":8080", s)
}
