package main

import (
	"fmt"
	"log"
	"net/http"
)

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (message String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, message)
}

func (message Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s%s %s", message.Greeting, message.Punct, message.Who)
}

func main() {
	// your http.Handle calls here
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
