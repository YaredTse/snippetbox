package main

import (
	"fmt"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Coming from parallel universe . . . "))
}

func main() {

	fmt.Println("Server started . . . . !")

	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)

	log.Println("Starting Server on Port :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
