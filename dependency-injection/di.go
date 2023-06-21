package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func GreetingHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "interwebbed Gophers")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(GreetingHandler)))
}
