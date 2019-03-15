package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	fmt.Println("Listening on port 8888...")
	router.HandleFunc("/", RootHandler).Methods("GET")
	router.HandleFunc("/isgomuxup", HealthCheck).Methods("GET")
	log.Fatal(http.ListenAndServe(":8888", router))
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Welcome to go-mux!")
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "go-mux is up!")
}
