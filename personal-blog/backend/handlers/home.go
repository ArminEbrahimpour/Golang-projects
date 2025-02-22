package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/ArminEbrahimpour/personal-blog/models"
	"github.com/julienschmidt/httprouter"
)

func HomePage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Implement your home page logic here
	articles, err := models.GetAllArticles()
	fmt.Println("All Artilcles successfully recived")
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to fetch articles", http.StatusInternalServerError)
		return
	}

	filePath := "../htmls/home.html"
	tmpl, err := template.ParseFiles(filePath)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to template", http.StatusInternalServerError)
	}

	sort.Slice(articles, func(i, j int) bool {
		date1, _ := time.Parse("2006-01-02", articles[i].Date)
		date2, _ := time.Parse("2006-01-02", articles[j].Date)
		return date1.After(date2)
	})
	// get 9 recent articles
	articles = articles[:9]

	data := struct {
		Title    string
		Articles []models.ArticleData
	}{
		Title:    "Personal Blog",
		Articles: articles,
	}

	// setting cookie to a normal user
	cookie := &http.Cookie{
		Name:     "cookie",
		Value:    "user",
		Expires:  time.Now().Add(24 * time.Hour),
		Path:     "/",
		HttpOnly: true,
	}
	// set the cookie on the response
	http.SetCookie(w, cookie)

	// before this code we should sort the articles based on the recent ones editted or created recently

	if err = tmpl.Execute(w, data); err != nil {
		log.Println(err)
		http.Error(w, "Unable to Execute the template", http.StatusInternalServerError)
	}

}
