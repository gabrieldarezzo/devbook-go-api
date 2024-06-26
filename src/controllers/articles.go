package controllers

import (
	response "api/src"
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	userId, erro := authentication.ExtractUserId(r)
	if erro != nil {
		response.ErroJSON(w, http.StatusUnauthorized, erro)
		return
	}

	bodyParams, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		response.ErroJSON(w, http.StatusBadRequest, erro)
		return
	}

	var article models.Article
	if erro = json.Unmarshal(bodyParams, &article); erro != nil {
		response.ErroJSON(w, http.StatusBadRequest, erro)
		return
	}
	article.AuthorId = userId

	if erro = article.Prepare(); erro != nil {
		response.ErroJSON(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connection()
	if erro != nil {
		response.ErroJSON(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryOfArticles(db)
	article.ID, erro = repository.CreateArticles(article)
	if erro != nil {
		response.ErroJSON(w, http.StatusBadRequest, erro)
		return
	}

	response.JSON(w, http.StatusOK, article)
}

func FindArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	articleId, erro := strconv.ParseUint(params["articleId"], 10, 64)
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

	repository := repositories.NewRepositoryOfArticles(db)
	article, erro := repository.FindArticle(articleId)
	if erro != nil {
		response.ErroJSON(w, http.StatusBadRequest, erro)
		return
	}

	response.JSON(w, http.StatusOK, article)
}

func FindArticles(w http.ResponseWriter, r *http.Request) {

	userId, erro := authentication.ExtractUserId(r)
	if erro != nil {
		response.ErroJSON(w, http.StatusUnauthorized, erro)
		return
	}

	db, erro := database.Connection()
	if erro != nil {
		response.ErroJSON(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryOfArticles(db)
	articles, erro := repository.FindArticles(userId)
	if erro != nil {
		response.ErroJSON(w, http.StatusBadRequest, erro)
		return
	}

	response.JSON(w, http.StatusOK, articles)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	articleId, erro := strconv.ParseUint(params["articleId"], 10, 64)
	if erro != nil {
		response.ErroJSON(w, http.StatusBadRequest, erro)
		return
	}

	userId, erro := authentication.ExtractUserId(r)
	if erro != nil {
		response.ErroJSON(w, http.StatusUnauthorized, erro)
		return
	}
	// fmt.Printf("\n%d -- DO user: userId: %d\n", articleId, userId)

	db, erro := database.Connection()
	if erro != nil {
		response.ErroJSON(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryOfArticles(db)
	articleInDb, erro := repository.FindArticle(articleId)
	if erro != nil {
		response.ErroJSON(w, http.StatusInternalServerError, erro)
		return
	}

	if userId != articleInDb.AuthorId {
		response.ErroJSON(w, http.StatusForbidden, errors.New("não é possível atualizar um post que não é seu"))
		return
	}

	bodyParams, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		response.ErroJSON(w, http.StatusBadRequest, erro)
		return
	}

	var articleToUpdate models.Article
	if erro = json.Unmarshal(bodyParams, &articleToUpdate); erro != nil {
		response.ErroJSON(w, http.StatusBadRequest, erro)
		return
	}

	if erro = articleToUpdate.Prepare(); erro != nil {
		response.ErroJSON(w, http.StatusBadRequest, erro)
		return
	}

	erro = repository.UpdateArticle(articleId, articleToUpdate)
	if erro != nil {
		response.ErroJSON(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	articleId, erro := strconv.ParseUint(params["articleId"], 10, 64)
	if erro != nil {
		response.ErroJSON(w, http.StatusBadRequest, erro)
		return
	}

	userId, erro := authentication.ExtractUserId(r)
	if erro != nil {
		response.ErroJSON(w, http.StatusUnauthorized, erro)
		return
	}

	db, erro := database.Connection()
	if erro != nil {
		response.ErroJSON(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryOfArticles(db)
	articleInDb, erro := repository.FindArticle(articleId)
	if erro != nil {
		response.ErroJSON(w, http.StatusInternalServerError, erro)
		return
	}

	if userId != articleInDb.AuthorId {
		response.ErroJSON(w, http.StatusForbidden, errors.New("não é possível excluir um post que não é seu"))
		return
	}

	erro = repository.DeleteArticle(articleId)
	if erro != nil {
		response.ErroJSON(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func GetAllArticlesFromUser(w http.ResponseWriter, r *http.Request) {
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

	repository := repositories.NewRepositoryOfArticles(db)
	articles, erro := repository.GetAllArticlesFromUser(userId)
	if erro != nil {
		response.ErroJSON(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, articles)
}

func IncreaseLikeInArticle(w http.ResponseWriter, r *http.Request) {
	articleId, erro := strconv.ParseUint(mux.Vars(r)["articleId"], 10, 64)
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

	repository := repositories.NewRepositoryOfArticles(db)
	erro = repository.IncreaseLikeInArticle(articleId)
	if erro != nil {
		response.ErroJSON(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, nil)
}
