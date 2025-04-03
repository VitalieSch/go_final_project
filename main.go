package main

import (
	"fmt"
	"net/http"
)

func main() {

	webDir := "web"

	http.Handle("/", http.FileServer(http.Dir(webDir)))

	fmt.Println("Server is listening...")

	err := http.ListenAndServe(":7540", nil)
	if err != nil {
		panic(err)
	}
}
