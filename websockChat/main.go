package main

import (
	"log"
	"net/http"
)

func serveIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, "templates/index.html")
}

func main() {

	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/websock", func(w http.ResponseWriter, r *http.Request) {
		serveHub(hub, w, r)
	})

	log.Fatal(http.ListenAndServe(":3000", nil))
}
