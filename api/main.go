package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Running API!")
	
	config.Load()
	fmt.Printf("Listening port: %d\n", config.Port)
	fmt.Println(config.DatabaseUrl)

	r := router.Generate()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
