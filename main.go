package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)


func main(){
	fmt.Println("Hello world")

	godotenv.Load(".env")
	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT is not found in the environment")
 	}

	router := chi.NewRouter()

	srv := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}

	err := srv.ListenAndServe()

	fmt.Printf("Server starting on port %v", portString)

	if err != nil {
		log.Fatal(err)
	}
}