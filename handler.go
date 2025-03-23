package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type ServerMessage struct{
	Code int `json:"code"`
	Message string `json:"message"`
}

func Handler() {

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
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func handleHomeRoute(w http.ResponseWriter, r *http.Request){
	var rootMessage ServerMessage = ServerMessage{Code: http.StatusOK, Message: "Server is functioning"}
	ipAddress := getIpAddr(r)
	if len(ipAddress) != 0{
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
	if len(ipAddress) != 0{
		fmt.Println(ipAddress)
	}
	fmt.Println(userInfo)
	json.NewEncoder(w).Encode(successMessage)

}

func getIpAddr(r *http.Request) []string {
	var ipAddr []string
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		ipAddr = append(ipAddr, fmt.Sprintf("xff %s", xff))
	}

	// Check X-Real-IP (another common header)
	xRealIP := r.Header.Get("X-Real-IP")
	if xRealIP != "" {
		ipAddr = append(ipAddr,fmt.Sprintf("xff %s", xRealIP))
	}

	// Fallback: use RemoteAddr
	ip := r.RemoteAddr
	ipAddr = append(ipAddr, fmt.Sprintf("xff %s", ip))
	return ipAddr
}