package main

import (
	"fmt"
	"go-chevere/config"
	"go-chevere/controller"
	"go-chevere/helper"
	"go-chevere/repository"
	"go-chevere/router"
	"go-chevere/service"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to the Go Server!")
// }

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

	db, err := config.ConnectDB()
	helper.ErrorPanic(err)

	defer db.Prisma.Disconnect()

	// repository
	postRepository := repository.NewPostRepository(db)

	// service
	postService := service.NewPostServiceImpl(postRepository)

	// controller
	postController := controller.NewPostController(postService)

	// router
	routes := router.NewRouter(postController)

	// mux := http.NewServeMux()
	// mux.HandleFunc("/", helloHandler)

	server := &http.Server{
		Addr:           ":" + port,
		Handler:        routes,
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
