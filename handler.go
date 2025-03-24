package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type ServerMessage struct{
	Code int `json:"code"`
	Message string `json:"message"`
}

func Handler(w http.ResponseWriter, r *http.Request) {

	err := godotenv.Load()
	if err != nil{
		fmt.Println("No ENV file present")
	}
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
	ipAddress := getIpAddr(r)
	if (ipAddress) != ""{
		fmt.Println(ipAddress)
	}
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
	ipAddress := getIpAddr(r)
	if (ipAddress) != ""{
		fmt.Println(ipAddress)
	}
	fmt.Println(userInfo)
	json.NewEncoder(w).Encode(successMessage)

}

func getIpAddr(r *http.Request) string {
	if vercelIP := r.Header.Get("X-Vercel-Forwarded-For"); vercelIP != "" {
        return vercelIP
    }
	
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		// The client IP is the first one in the list
		ips := strings.Split(xff, ",")
		if len(ips) > 0 {
			return strings.TrimSpace(ips[0]) // Return the first IP
		}
	}

	// Fallback to X-Real-IP
	xRealIP := r.Header.Get("X-Real-IP")
	if xRealIP != "" {
		return xRealIP
	}

	// Final fallback: use RemoteAddr
	return r.RemoteAddr
}