package handlers

import (
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/ArminEbrahimpour/personal-blog/models"
	"github.com/julienschmidt/httprouter"
)

func AdminPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Implement your admin page logic here
	articles, err := models.GetAllArticles()
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to get the articles ", http.StatusInternalServerError)
	}
	filePath := "../htmls/admin.html"
	tmpl, err := template.ParseFiles(filePath)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to load the template", http.StatusInternalServerError)
	}
	data := struct {
		Title    string
		Articles []models.ArticleData
	}{
		Title:    "Personal Blog",
		Articles: articles,
	}

	if err = tmpl.Execute(w, data); err != nil {
		log.Println(err)
		http.Error(w, " Unable to Execute the template", http.StatusInternalServerError)
	}

}

func Protect(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Implement your protection logic here
		cookie, err := r.Cookie("cookie")
		if err != nil {
			log.Println(err)
			http.Error(w, "cookie didn't attache", http.StatusInternalServerError)

		}
		if cookie.Value == "Admin" {
			h(w, r, ps)
		} else {

			cookie := &http.Cookie{
				Name:     "cookie",
				Value:    "user",
				Expires:  time.Now().Add(24 * time.Hour),
				Path:     "/",
				HttpOnly: true,
			}
			// set the cookie on the response
			http.SetCookie(w, cookie)

			w.Write([]byte("You Don't Have access to this page(get the fuck out)"))

		}
	}
}
