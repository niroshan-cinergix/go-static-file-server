package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./pages")))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}