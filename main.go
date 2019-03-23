package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Response struct {
	Status int    `json:status`
	Body   string `json:body`
}

func main() {
	router := mux.NewRouter()
	fmt.Println("Listening on port 8888...")
	router.HandleFunc("/", RootHandler).Methods("GET")
	router.HandleFunc("/isgomuxup", HealthCheck).Methods("GET")
	log.Fatal(http.ListenAndServe(":8888", router))
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	response := &Response{
		Status: http.StatusOK,
		Body:   "Welcome to go-mux!",
	}
	data, _ := json.Marshal(response)
	w.WriteHeader(response.Status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	//json.NewEncoder(w).Encode("Welcome to go-mux!")
	//fmt.Fprint(w, "Welcome to go-mux!")
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := &Response{
		Status: http.StatusOK,
		Body:   "go-mux is up!",
	}
	data, _ := json.Marshal(response)
	w.WriteHeader(response.Status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	//json.NewEncoder(w).Encode("Welcome to go-mux!")
	//fmt.Fprint(w, "go-mux is up!")
}
