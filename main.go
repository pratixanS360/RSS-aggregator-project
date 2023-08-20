package main

import (
       "fmt"
       "log"
       "os"
	   "net/http"
	   "github.com/joho/godotenv"
	   "github.com/go-chi/chi"
	   //"github.com/go-chi/cors"
)

func main() {
	fmt.Println("Entry point to the RSS aggregator!")

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

	log.Printf("Server starting on PORT %v", portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("PORT: ", portString)
}