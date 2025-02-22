package main

import (
	"net/http"

	"github.com/ArminEbrahimpour/personal-blog/router"
)

func main() {
	router := router.NewRouter()

	server := http.Server{
		Addr:    "127.0.0.1:15263",
		Handler: router,
	}

	server.ListenAndServe()
}
