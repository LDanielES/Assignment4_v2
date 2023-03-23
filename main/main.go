package main

import (
	"Assinment4_Copy/Client"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Parse HTML file
var HtmlParser = template.Must(template.ParseFiles("/Frontend/index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	buf := &bytes.Buffer{}
	err := HtmlParser.Execute(buf, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	buf.WriteTo(w)
}

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
	StaticFiles := http.FileServer(http.Dir("Frontend"))
	mux.Handle("/Frontend/", http.StripPrefix("/Frontend/", StaticFiles))

	// Start server
	mux.HandleFunc("/search", Search.searchHandler(newsApi))
	mux.HandleFunc("/", indexHandler)

	fmt.Println("Listening on :8088...")
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
