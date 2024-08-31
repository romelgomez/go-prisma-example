package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Go Server!")
}

func main() {
	fmt.Println("Start Server")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to port 8080 if not specified
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)

	server := &http.Server{
		Addr:           ":" + port,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("Server running on port %s\n", port)
	server_err := server.ListenAndServe()

	if server_err != nil {
		panic(server_err)
	}
}
