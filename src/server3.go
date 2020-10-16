package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/gendata", gendata)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}

//!+handler
// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {

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
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}

}

func gendata(w http.ResponseWriter, r *http.Request) {
	numBytes, err := r.URL.Query()["numBytes"]

	if err != true {
		log.Print(err)
	}
	if len(numBytes[0]) < 1 {
		fmt.Fprintf(w, "No value for \"numBytes\" was given\n")
		fmt.Fprintf(w, "You can add the value to the end of the URL via the format:\n")
		fmt.Fprintf(w, "\t/gendata?numBytes=<value>")
		return
	}

	byteCount, err1 := strconv.Atoi(numBytes[0])
	if err1 != nil {
		log.Print(err)
	}

	result := strings.Repeat("q", byteCount)
	fmt.Fprintf(w, "%s", result)
}

//!-handler
