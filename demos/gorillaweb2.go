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

// GoodNightHandler tells the Gorilla goodnight.
func GoodNightHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Say goodnight Gracie!\n"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/goodnight", GoodnightHandler)
	r.HandleFunc("/", GorillaHandler)

	log.Println("Started server on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}
