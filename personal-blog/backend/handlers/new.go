package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/ArminEbrahimpour/personal-blog/models"
	"github.com/julienschmidt/httprouter"
)

func GetNewPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	data := struct {
		Title string
	}{
		Title: "New Article",
	}
	// omel
	filePath := "../htmls/NewArticle.html"
	tmpl, err := template.ParseFiles(filePath)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to load the html", http.StatusInternalServerError)
	}
	if err = tmpl.Execute(w, data); err != nil {
		log.Println(err)
		http.Error(w, "Unalbe to execute the template", http.StatusInternalServerError)
	}

}

func NewArticle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var article models.Article
	// creating an Article typed object
	var articleData models.ArticleData

	// get the json from the body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, "Couldn't read the request body", http.StatusInternalServerError)
	}

	if err = json.Unmarshal(body, &article); err != nil {
		log.Println(err)
		http.Error(w, "Couldn't unmarshal the request", http.StatusInternalServerError)
	}

	articles, err := models.GetAllArticles()
	if err != nil {
		log.Println(err)
		http.Error(w, "Couldn't retrive all the articles ", http.StatusInternalServerError)
	}
	// converting the article to ArticleData type
	articleData.Id = len(articles) + 1
	articleData.Title = article.Title
	articleData.Heading = article.Heading
	articleData.Content = article.Content
	articleData.Date = article.Date

	articles = append(articles, articleData) // now we have added our new article

	var wrapper models.ArticlesWrapper
	wrapper.Articles = articles

	// creatin a file for new articles.json file

	f, err := os.Create("../storage/articles.json")
	if err != nil {
		log.Println(err)
		http.Error(w, "couldn't create the file", http.StatusInternalServerError)
	}
	defer f.Close()
	if err = json.NewEncoder(f).Encode(&wrapper); err != nil {
		log.Println(err)
		http.Error(w, "couldn't encode the wrapper", http.StatusInternalServerError)
	}
}
