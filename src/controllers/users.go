package controllers

import (
	"api/database"
	"api/repositories"
	response "api/src"
	"api/src/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// CreateUser Create a User
func CreateUser(w http.ResponseWriter, r *http.Request) {

	bodyRequest, erro := ioutil.ReadAll(r.Body)
	// fmt.Println("Corpo da solicitação:", string(bodyRequest))

	if erro != nil {
		response.ErroJSON(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User
	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		response.ErroJSON(w, http.StatusBadRequest, erro)
	}

	if erro = user.Prepare(); erro != nil {
		response.ErroJSON(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connection()
	if erro != nil {
		response.ErroJSON(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositoryUser := repositories.NewRepositoryOfUsers(db)
	user.ID, erro = repositoryUser.CreateUser(user)
	user.Password = "***"
	if erro != nil {
		response.ErroJSON(w, http.StatusInternalServerError, erro)
	}
	response.JSON(w, http.StatusCreated, user)
}

// FindUsers Encontra um
func FindUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os Usuário!"))
}

// FindUser(w ht
func FindUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um Usuário!"))
}

// UpdateUser(w
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando Usuário!"))
}

// DeleteUser(w
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando Usuário!"))
}
