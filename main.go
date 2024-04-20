package main

import (
	"api/sql/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.SetEnvs()
	fmt.Printf("\nRunning API in port: %d!\n", config.Port)

	r := router.Generate()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))

}
