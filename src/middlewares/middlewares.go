package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

// Logger write information of requests in terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("Autenticando...")
		log.Printf("\n [%s] %s %s \n", r.Method, r.URL, r.Host)
		next(w, r)
	}
}

// Authenticate check if user is authenticated
func Authenticate(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Autenticando...")
		next(w, r)
	}
}
