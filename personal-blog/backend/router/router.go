package router

import (
	"github.com/julienschmidt/httprouter"

	"github.com/ArminEbrahimpour/personal-blog/handlers"
)

func NewRouter() *httprouter.Router {

	router := httprouter.New()

	router.GET("/home", handlers.HomePage)
	router.GET("/admin", handlers.Protect(handlers.AdminPage))
	router.GET("/new", handlers.Protect(handlers.GetNewPage))
	router.POST("/new", handlers.Protect(handlers.NewArticle))
	router.GET("/article/:number", handlers.Protect(handlers.ShowArticle))
	router.GET("/edit/:number", handlers.Protect(handlers.GetEditPage))
	router.POST("/edit/:number", handlers.Protect(handlers.EditArticle))
	router.GET("/delete/:number", handlers.Protect(handlers.DeleteArticle))
	router.GET("/", handlers.HomePage)

	return router

}
