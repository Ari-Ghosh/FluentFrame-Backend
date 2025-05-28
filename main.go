package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type Message struct {
	Text     string `json:"text"`
	DateTime string `json:"datetime"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// API endpoint
	http.HandleFunc("/api/message", func(w http.ResponseWriter, r *http.Request) {
		// Enable CORS for development
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		response := Message{
			Text:     "Hello from Go!",
			DateTime: time.Now().Format(time.RFC3339),
		}

		json.NewEncoder(w).Encode(response)
	})

	// Serve static files (for production)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}