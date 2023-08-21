package main

import (
       "fmt"
       "log"
       "os"
	   "net/http"
	   "github.com/joho/godotenv"
	   "github.com/go-chi/chi"
	   "github.com/go-chi/cors"
)

func main() {
	fmt.Println("Entry point to the RSS aggregator!")

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
	   log.Fatal("PORT is not found in the environment")
	}

	router := chi.NewRouter()

	// Middleware for handling CORS
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		//ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	srv := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}

	//Todo: Understand??
	router.Get("", func(w srv.ResponseWriter, r *http.Request)) {
		w.write([]byte("API Response"))
	}

	log.Printf("Server starting on PORT %v", portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("PORT: ", portString)
}