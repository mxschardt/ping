package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	PORT    = "3000"
	MESSAGE = "default"
)

func init() {
	port, exists := os.LookupEnv("PORT")
	if exists {
		PORT = port
	}

	message, exists := os.LookupEnv("MESSAGE")
	if exists {
		MESSAGE = message
	}
}

func main() {
	fmt.Printf("serving on port %s", PORT)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+PORT, nil))
}

// handler echoes the Path component of the requested URL.
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

	fmt.Fprintf(w, "\n%s", MESSAGE)
}
