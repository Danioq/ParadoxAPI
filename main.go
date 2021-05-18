package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Danioq/ParadoxAPI/http"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	listen := os.Getenv("LISTEN")

	server := http.GetServer(listen)
	fmt.Println("Listen on:", listen)
	err = server.ListenAndServe()

	if err != nil {
		fmt.Println(err)
		panic(1)
	}

}
