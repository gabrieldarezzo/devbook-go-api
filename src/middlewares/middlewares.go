package middlewares

import (
	response "api/src"
	"api/src/authentication"
	"log"
	"net/http"
)

// Logger write information of requests in terminal
func Logger(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("Autenticando...")
		log.Printf("\n [%s] %s %s \n", r.Method, r.URL, r.Host)
		nextFunction(w, r)
	}
}

// Authenticate check if user is authenticated
func Authenticate(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := authentication.ValidateToken(r); erro != nil {
			response.ErroJSON(w, http.StatusUnauthorized, erro)
			return
		}

		nextFunction(w, r)
	}
}
