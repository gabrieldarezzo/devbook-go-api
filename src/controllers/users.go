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
		return
	}

	if erro = user.Prepare("NEW_USER"); erro != nil {
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

	if user.ID == 0 {
		response.JSON(w, http.StatusNotFound, nil)
		return
	}

	response.JSON(w, http.StatusOK, user)
}

// FindUser Find one or more user using criteria
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	userId, erro := strconv.ParseUint(params["userId"], 10, 64)
	if erro != nil {
		response.ErroJSON(w, http.StatusBadRequest, erro)
		return
	}

	userIdToken, erro := authentication.ExtractUserId(r)
	if erro != nil {
		response.ErroJSON(w, http.StatusUnauthorized, erro)
		return
	}

	if userIdToken != userId {
		response.ErroJSON(w, http.StatusForbidden, errors.New("Não é possivel atualizar um usuario que não é seu"))
		return
	}

	userBodyParams, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		response.ErroJSON(w, http.StatusBadRequest, erro)
		return
	}

	var userToUpdate models.User
	if erro = json.Unmarshal(userBodyParams, &userToUpdate); erro != nil {
		response.ErroJSON(w, http.StatusBadRequest, erro)
		return
	}

	if erro = userToUpdate.Prepare("UPDATE_USER"); erro != nil {
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
	erro = repositoryUser.UpdateUserById(userId, userToUpdate)
	if erro != nil {
		response.ErroJSON(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

// DeleteUser
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, erro := strconv.ParseUint(params["userId"], 10, 64)
	if erro != nil {
		response.ErroJSON(w, http.StatusBadRequest, erro)
		return
	}

	userIdToken, erro := authentication.ExtractUserId(r)
	if erro != nil {
		response.ErroJSON(w, http.StatusUnauthorized, erro)
		return
	}

	if userIdToken != userId {
		response.ErroJSON(w, http.StatusForbidden, errors.New("Não é possivel atualizar um usuario que não é seu"))
		return
	}

	db, erro := database.Connection()
	if erro != nil {
		response.ErroJSON(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositoryUser := repositories.NewRepositoryOfUsers(db)
	erro = repositoryUser.DeleteUser(userId)
	if erro != nil {
		response.ErroJSON(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

// FallowUser Fallow another user
func FollowUser(w http.ResponseWriter, r *http.Request) {

	userIdToken, erro := authentication.ExtractUserId(r)
	if erro != nil {
		response.ErroJSON(w, http.StatusUnauthorized, erro)
		return
	}

	params := mux.Vars(r)

	userId, erro := strconv.ParseUint(params["userId"], 10, 64)

	if erro != nil {
		response.ErroJSON(w, http.StatusBadRequest, erro)
		return
	}

	if userIdToken == userId {
		response.ErroJSON(w, http.StatusForbidden, errors.New("não é possível seguir você mesmo"))
		return
	}

	db, erro := database.Connection()
	if erro != nil {
		response.ErroJSON(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositoryUser := repositories.NewRepositoryOfUsers(db)
	erro = repositoryUser.FallowNewUser(userIdToken, userId)
	if erro != nil {
		response.ErroJSON(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusCreated, nil)

}

// UnFollowUser UnFollowUser a user
func UnFollowUser(w http.ResponseWriter, r *http.Request) {

	userIdToken, erro := authentication.ExtractUserId(r)
	if erro != nil {
		response.ErroJSON(w, http.StatusUnauthorized, erro)
		return
	}

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
	erro = repositoryUser.UnFollowUser(userIdToken, userId)
	if erro != nil {
		response.ErroJSON(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)

}

// GetFollowers get all fallowers of specific user
func GetFollowers(w http.ResponseWriter, r *http.Request) {
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
	followers, erro := repositoryUser.FollowersOfByUserId(userId)
	if erro != nil {
		response.ErroJSON(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, followers)

}

// GetFollowing get all following users of specific user
func GetFollowing(w http.ResponseWriter, r *http.Request) {
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
	followingsUsers, erro := repositoryUser.GetAllFollowingUsersOfUserId(userId)
	if erro != nil {
		response.ErroJSON(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, followingsUsers)
}

// UpdatePassword Update a password of own user
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	userIdToken, erro := authentication.ExtractUserId(r)
	if erro != nil {
		response.ErroJSON(w, http.StatusUnauthorized, erro)
		return
	}
	params := mux.Vars(r)

	userId, erro := strconv.ParseUint(params["userId"], 10, 64)
	if erro != nil {
		response.ErroJSON(w, http.StatusBadRequest, erro)
		return
	}

	if userIdToken != userId {
		response.ErroJSON(w, http.StatusForbidden, errors.New("não é possível trocar a senha de outra pessoa"))
		return
	}

	bodyParams, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		response.ErroJSON(w, http.StatusBadRequest, erro)
		return
	}

	var password models.Password
	if erro = json.Unmarshal(bodyParams, &password); erro != nil {
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
	passwordHashedInDb, erro := repositoryUser.GetPaswordByIdUser(userId)

	if erro = security.CheckPassword(passwordHashedInDb, password.ActualPassword); erro != nil {
		response.ErroJSON(w, http.StatusBadRequest, errors.New("Senha incorreta"))
		return
	}

	newPasswordHashed, erro := security.GenerateHash(password.NewPasword)

	if erro != nil {
		response.ErroJSON(w, http.StatusBadRequest, errors.New("Senha incorreta"))
		return
	}

	erro = repositoryUser.UpdatePasswordUserById(userId, string(newPasswordHashed))
	if erro != nil {
		response.ErroJSON(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
