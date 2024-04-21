package main

import (
	"api/sql/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

// func init() {
// 	keyString := make([]byte, 64)

// 	if _, erro := rand.Read(keyString); erro != nil {
// 		log.Fatal(erro)
// 	}

// 	stringBase64 := base64.StdEncoding.EncodeToString(keyString)
// 	fmt.Println(stringBase64)
// }

func main() {
	config.SetEnvs()
	fmt.Printf("\nRunning API in port: %d!\n", config.Port)
	r := router.Generate()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))

}
