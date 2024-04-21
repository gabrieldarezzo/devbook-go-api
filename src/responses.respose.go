package response

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
)

// JSON return response in JSON format
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data == nil {
		w.Write([]byte("{}"))
		return
	}

	if reflect.TypeOf(data).Kind() == reflect.Slice {
		length := reflect.ValueOf(data).Len()
		if length == 0 {

			w.Write([]byte("{}"))
			return
		}
	}

	if erro := json.NewEncoder(w).Encode(data); erro != nil {
		log.Fatal(erro)
	}
}

// ErroJSON return response in Json format
func ErroJSON(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}
