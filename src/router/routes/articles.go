package routes

import (
	"api/src/controllers"
	"net/http"
)

var routeArticles = []Route{
	{
		URI:                    "/articles",
		Method:                 http.MethodPost,
		Function:               controllers.CreateArticle,
		RequiredAuthentication: true,
	},
	{
		URI:                    "/articles",
		Method:                 http.MethodGet,
		Function:               controllers.FindArticles,
		RequiredAuthentication: true,
	},
	{
		URI:                    "/articles/{articleId}",
		Method:                 http.MethodGet,
		Function:               controllers.FindArticle,
		RequiredAuthentication: true,
	},
	{
		URI:                    "/articles/{articleId}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdateArticle,
		RequiredAuthentication: true,
	},
	{
		URI:                    "/articles/{articleId}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeleteArticle,
		RequiredAuthentication: true,
	},
}
