package main

import (
       "fmt"
       "log"
       "os"
	   "github.com/joho/godotenv"
)

func main() {
	fmt.Println("Entry point to the RSS aggregator!")

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
	   log.Fatal("PORT is not found in the environment")
	}

	fmt.Println("PORT: ", portString)
}