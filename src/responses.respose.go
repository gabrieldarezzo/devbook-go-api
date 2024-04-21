package response

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
)

// JSON return response in JSON format
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	// Define o código de status da resposta
	w.WriteHeader(statusCode)

	// Se os dados forem nil, envie uma resposta vazia
	if data == nil {
		return
	}

	// Se os dados forem um slice vazio, envie um JSON vazio
	if reflect.TypeOf(data).Kind() == reflect.Slice {
		sliceValue := reflect.ValueOf(data)
		if sliceValue.Len() == 0 {
			// Envia um JSON vazio
			_, err := w.Write([]byte("{}"))
			if err != nil {
				log.Fatal(err)
			}
			return
		}
	}

	// Se não for um slice vazio e não for nil, codifique os dados para JSON e envie-os na resposta
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
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
