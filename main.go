package main

import (
	"log"
	"net/http"
)

func main() {
	setupAPI()

	// Serve on port :8000, fudge yeah hardcoded port
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// setupAPI will start all Routes and their Handlers
func setupAPI() {
	// Create a Manager instance used to handle WebSocket Connections
	manager := NewManager()

	// Serve the ./public directory at Route /
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/ws", manager.serveWS)
}
