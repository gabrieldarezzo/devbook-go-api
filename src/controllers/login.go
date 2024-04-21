package controllers

import (
	response "api/src"
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/security"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// CreateUser Create a User
func Login(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		response.ErroJSON(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User
	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
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
	userInDatabase, erro := repositoryUser.FindUserByEmail(user.Email)

	if userInDatabase.ID == 0 {
		response.ErroJSON(w, http.StatusBadRequest, errors.New("email ou senha não encontrado"))
		return
	}

	if erro = security.CheckPassword(userInDatabase.Password, user.Password); erro != nil {
		// response.ErroJSON(w, http.StatusUnauthorized, erro)
		response.ErroJSON(w, http.StatusBadRequest, errors.New("email ou senha não encontrado"))
		return
	}

	token, _ := authentication.CreateToken(userInDatabase.ID)

	response.JSON(w, http.StatusOK, token)

}
