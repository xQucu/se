package main

import (
	"log"
	"net/http"
	"stayease-app"
)

func main() {
	server := stayease.NewServer()
	log.Println("StayEase role-based local server starting on :8080...")
	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
