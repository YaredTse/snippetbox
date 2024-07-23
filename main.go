package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Coming from parallel universe . . . "))
}

func snippetView(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	msg := fmt.Sprintf("Display a specific nippet with ID %d ", id)
	w.Write([]byte(msg))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form form for creating a new snippet . . . . "))
}

func main() {

	fmt.Println("Server started . . . . !")

	mux := http.NewServeMux()
	mux.HandleFunc("apix.apigee.com/{$}", home)
	mux.HandleFunc("/snippet/view/{id}", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting Server on Port :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
