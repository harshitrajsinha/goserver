package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type ServerMessage struct{
	Code int `json:"code"`
	Message string `json:"message"`
}

func Handler(w http.ResponseWriter, r *http.Request) {

	_ = godotenv.Load()
	
	port := os.Getenv("PORT")
	if port == ""{
		port = "8080"
	}

	// Create router using gorilla/mux
	var router *mux.Router = mux.NewRouter()

	// Create routes
	router.HandleFunc("/", handleHomeRoute).Methods("GET")
	router.HandleFunc("/info", handleUserInformation).Methods("POST")

	// create server
	fmt.Println("Listening at PORT ", port)
	router.ServeHTTP(w, r)
}

func handleHomeRoute(w http.ResponseWriter, r *http.Request){
	var rootMessage ServerMessage = ServerMessage{Code: http.StatusOK, Message: "Server is functioning"}
	getIpAddr(r)

	json.NewEncoder(w).Encode(rootMessage)
}

func handleUserInformation(w http.ResponseWriter, r *http.Request){

	var userInfo map[string]interface{}
	var successMessage ServerMessage = ServerMessage{Code: http.StatusOK, Message: "Information recieved successfully"}

	// get user information from request body
	err := json.NewDecoder(r.Body).Decode(&userInfo)
	if err != nil{
		http.Error(w, "Invalid format", http.StatusBadRequest)
		fmt.Println("Error occured ", err)
		return
	}
	getIpAddr(r)

	fmt.Println(userInfo)
	json.NewEncoder(w).Encode(successMessage)

}

func getIpAddr(r *http.Request) {

	vercelIP := r.Header.Get("X-Vercel-Forwarded-For")
	fmt.Println("vercelip, ", vercelIP)
	
	xff := r.Header.Get("X-Forwarded-For")
	fmt.Println("xff, ", xff)

	// Fallback to X-Real-IP
	xRealIP := r.Header.Get("X-Real-IP")
	fmt.Println("xRealIP, ", xRealIP)

	// Final fallback: use RemoteAddr
	fmt.Println("r.RemoteAddr ", r.RemoteAddr)
}