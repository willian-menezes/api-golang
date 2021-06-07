package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()
	fmt.Println("API is running...")

	r := router.Make()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
