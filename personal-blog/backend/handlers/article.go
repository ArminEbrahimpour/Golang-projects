package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/ArminEbrahimpour/personal-blog/models"

	"github.com/julienschmidt/httprouter"
)

func ShowArticle(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	articleNumber, err := strconv.Atoi(p.ByName("number"))
	if err != nil {

		http.Error(w, "Invalid article number ", http.StatusBadRequest)

	}
	article, err := models.FetchArticleData(articleNumber)
	fmt.Println("function called and this is the article")
	fmt.Println(article)
	if err != nil {
		log.Println(err)

		http.Error(w, "Invalid article number ", http.StatusBadRequest)

	}
	filePath := "../htmls/article.html"

	tmpl, err := template.ParseFiles(filePath)

	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid article number ", http.StatusBadRequest)
	}

	if err = tmpl.Execute(w, article); err != nil {
		log.Println(err)
		http.Error(w, "Invalid article number ", http.StatusBadRequest)
	}
}
