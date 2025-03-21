package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ServerMessage struct{
	Message string `json:"message"`
}

func main(){

	router := mux.NewRouter()

	router.HandleFunc("/", root).Methods("GET")

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func root(w http.ResponseWriter, r *http.Request){
	messageJson := ServerMessage{Message:"Server is functioning"}
	json.NewEncoder(w).Encode(messageJson)
}