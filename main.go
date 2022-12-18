package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hiteshshimpi-55/go-mongo-api/router"
)

func main() {
	fmt.Println("Hello Bitches...")
	r := router.Router()
	fmt.Println("Server started at port 8080 . . . ")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println("Error While listening")
		log.Fatal(err)
	}
}
