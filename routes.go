package main

import (
	"github.com/julienschmidt/httprouter"
	"github.corp.globant.com/a-muliarchik/GoTraining/bookPackage"
)

// NewRouter return all router
func NewRouter(handle *bookPackage.BaseHandler) *httprouter.Router {

	router := httprouter.New()
	router.GET("/", handle.Index)
	router.GET("/books", handle.GetAllBooks)
	router.POST("/books", handle.CreateBook)
	router.GET("/books/:id", handle.GetBookByID)
	router.PUT("/books/:id", handle.UpdateBook)
	router.DELETE("/books/:id", handle.DeleteBook)
	return router
}