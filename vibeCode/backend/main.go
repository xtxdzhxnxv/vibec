package main

import (
	"log"
	"net/http"
	"os"
	"vibecode/backend/handlers"
	"vibecode/backend/database"
)

func main() {
	database.InitDB()
	log.Println("Backend server started successfully.")

	http.HandleFunc("/api/register", (handlers.Register))
	http.HandleFunc("/api/login", (handlers.Login))
	http.HandleFunc("/api/deleteaccount", (handlers.DeleteAccount))
	http.HandleFunc("/api/updateaccount", (handlers.UpdateAccount))
	http.HandleFunc("/api/getaccounts", (handlers.GetAccounts))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Backend server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
