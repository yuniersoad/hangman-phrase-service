package main

import (
	"fmt"
	_ "github.com/yuniersoad/hangman-phrase-service/handlers"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
