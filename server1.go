package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/product", handler)
	http.HandleFunc("/product/", handler2)
	http.HandleFunc("/product/sort", sortHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!+handler
// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "handler1\n")
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "%q = %q\n", k, v)
	}
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "handler2\n")
	fmt.Fprintf(w, "Method: %q\n", r.Method)
	var path = strings.Split(r.URL.Path, "/")
	fmt.Fprintln(w, "path", path)
}

func sortHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Sorted:")
}
