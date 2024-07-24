package templates

func RenderHttpTemplate() string {
	const mainTemplate = `package main

import (
	"fmt"
	"io"
	"net/http"
)

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func main() {
	http.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Println(err)
	}
}
	`

	return mainTemplate
}
