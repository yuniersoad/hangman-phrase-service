package main

import (
	"fmt"
	_ "github.com/yuniersoad/hangman-phrase-service/handlers"
	"github.com/yuniersoad/hangman-phrase-service/storage"
	"log"
	"net/http"
	"os"
)

func main() {
	dbHost, ok := os.LookupEnv("DB_HOST")
	if !ok {
		dbHost = "127.0.0.1"
	}
	err := storage.Setup(dbHost)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
