package main

import (
	"api/src/config"
	"api/src/router"
	// "crypto/rand"
	// "encoding/base64"
	"fmt"
	"log"
	"net/http"
)

// Used only once, generated secret key
// func init() {
// 	key := make([]byte, 64)

// 	if _, err := rand.Read(key); err != nil {
// 		log.Fatal(err)
// 	}

// 	base64Key := base64.StdEncoding.EncodeToString(key)
// 	fmt.Println(base64Key)	
// }

func main() {
	fmt.Println("Running API!")
	
	config.Load()
	fmt.Printf("Listening port: %d\n", config.Port)
	fmt.Println(config.DatabaseUrl)
	fmt.Println(config.SecretKey)

	r := router.Generate()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
