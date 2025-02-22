package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"

	"github.com/ArminEbrahimpour/personal-blog/models"
	"github.com/julienschmidt/httprouter"
)

func GetEditPage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	data := struct {
		Title string
	}{
		Title: "Edit Article",
	}
	filePath := "../htmls/NewArticle.html"
	tmpl, err := template.ParseFiles(filePath)
	if err != nil {
		log.Println(err)
	}
	if err = tmpl.Execute(w, data); err != nil {
		log.Println(err)
	}
}
func EditArticle(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Implement your edit article logic here
	var article models.Article
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	if err = json.Unmarshal(body, &article); err != nil {
		log.Println(err)
	}
	DeleteArticle(w, r, p)
	articles, err := models.GetAllArticles()

	if err != nil {
		log.Println(err)
	}
	var articleData models.ArticleData
	articleData.Id = len(articles) + 1
	articleData.Title = article.Title
	articleData.Heading = article.Heading
	articleData.Content = article.Content
	articleData.Date = article.Date

	var wrapper models.ArticlesWrapper

	articles = append(articles, articleData)

	wrapper.Articles = articles

	f, err := os.Create("../storage/articles.json")
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	if err = json.NewEncoder(f).Encode(&wrapper); err != nil {
		log.Println(err)
	}
}

func DeleteArticle(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// delete article function

	// get all articles
	articles, err := models.GetAllArticles()
	if err != nil {
		log.Println(err)
	}

	// put them in an slice
	num, err := strconv.Atoi(p.ByName("number"))

	var newArticles []models.ArticleData

	//index := 0
	// create a new slice with [:p] and [p+1:]
	for _, article := range articles {
		if article.Id != num {
			if article.Id > num {
				article.Id -= 1
			}
			newArticles = append(newArticles, article)
		}
	}
	// using os lib create new article.json file
	var wrapper models.ArticlesWrapper
	wrapper.Articles = newArticles

	f, err := os.Create("../storage/articles.json")
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if err = json.NewEncoder(f).Encode(&wrapper); err != nil {
		log.Println(err)
	}
}
