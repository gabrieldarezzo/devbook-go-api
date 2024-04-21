package controllers

import (
	"api/database"
	"api/repositories"
	response "api/src"
	"api/src/models"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CreateUser Create a User
func CreateUser(w http.ResponseWriter, r *http.Request) {

	bodyRequest, erro := ioutil.ReadAll(r.Body)

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
	if erro != nil {
		response.ErroJSON(w, http.StatusInternalServerError, erro)
		return
	}
	user.Password = "***"
	response.JSON(w, http.StatusCreated, user)
}

// FindUsers Show a couple of users
func FindUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))
	nameOrNick = strings.TrimSpace(nameOrNick)

	if nameOrNick == "" {
		response.ErroJSON(w, http.StatusInternalServerError, errors.New("O queryParam: 'user' é obrigatório e não pode estar em branco"))
		return
	}

	db, erro := database.Connection()
	if erro != nil {
		response.ErroJSON(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositoryUser := repositories.NewRepositoryOfUsers(db)
	users, erro := repositoryUser.FindUser(nameOrNick)
	if erro != nil {
		response.ErroJSON(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, users)
}

// FindUser Find one or more user using criteria
func FindUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	userId, erro := strconv.ParseUint(params["userId"], 10, 64)

	if erro != nil {
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
	user, erro := repositoryUser.FindUserById(userId)
	if erro != nil {
		response.ErroJSON(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, user)
}

// UpdateUser Update User
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando Usuário!"))
}

// DeleteUser
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando Usuário!"))
}
