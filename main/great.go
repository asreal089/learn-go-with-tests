package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Great(wr io.Writer, name string) {
	fmt.Fprintf(wr, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Great(w, "World")
}

// change to main to expose endpoint.
func mainGreat() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
