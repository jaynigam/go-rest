//This is the first way to achieve REST API in GoLang

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Rest Api has started")
	handleRequests()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	log.Println("Welcome to Homepage")
	fmt.Println("Welcome to our Homepage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Println("Request is available")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
