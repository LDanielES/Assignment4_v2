package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/LDanielES/Assinment4_Copy/Client"
	HP "github.com/LDanielES/Assinment4_Copy/HtmlParser"
	"github.com/LDanielES/Assinment4_Copy/Search"
	"github.com/joho/godotenv"
)

func main() {
	//Create HTTP request multiplexer
	mux := http.NewServeMux()

	//Load environment variables from .env file in root directory
	err := godotenv.Load("build.env")
	if err != nil {
		log.Println("Error loading .env file")
	}
	//Get PORT value
	port := os.Getenv("PORT")

	//Retrieve the News API key from the environment
	apiKey := os.Getenv("NEWS_API_KEY")
	if apiKey == "" {
		log.Fatal("Env: apiKey must be set")
	}

	//Create a new Client instance
	myClient := &http.Client{Timeout: 10 * time.Second}
	newsApi := Client.NewClient(myClient, apiKey, 20)

	//Static served files
	StaticFiles := http.FileServer(http.Dir("style"))
	mux.Handle("/style/", http.StripPrefix("/style/", StaticFiles))

	// Start server
	mux.HandleFunc("/search", Search.SearchHandler(newsApi))
	mux.HandleFunc("/", HP.IndexHandler)

	fmt.Println("Listening on :8088...")
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
