package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// GorillaHandler shouts to wake up the gorilla.
func GorillaHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hey Gorilla!\n"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", GorillaHandler)

	log.Println("Started server on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}
