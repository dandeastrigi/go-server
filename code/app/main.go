package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// main function
	http.HandleFunc("/", RequestHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Service Started on Port " + port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

// RequestHandler : handle requests
func RequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, `{"status":"ok"}`)
}
