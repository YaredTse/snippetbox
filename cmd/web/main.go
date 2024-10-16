package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {

	// DEfine a new command-line flag with name addr a default value of :4000
	// And some short help explaining what the flag controls.  The value of the flag
	// will be stored in the addr variable at runtime
	addr := flag.String("addr", ":4000", "HTTP Network address")

	// Importantly we use the flag.Parse()
	flag.Parse()

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)
	mux.HandleFunc("GET /testtest", home)

	log.Printf("starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)

	log.Fatal(err)
}
