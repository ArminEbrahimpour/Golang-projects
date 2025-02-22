package models

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type ArticleData struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Heading string `json:"heading"`
	Content string `json:"content"`
	Date    string `json:"date"`
}

type ArticlesWrapper struct {
	Articles []ArticleData `json:"articles"`
}

type Article struct {
	Title   string `json:"title"`
	Heading string `json:"heading"`
	Content string `json:"content"`
	Date    string `json:"date"`
}

/*
fetching data from a file called articles.json and getting the article with specified Id
*/
func FetchArticleData(articleNumber int) (ArticleData, error) {
	var wrapper ArticlesWrapper

	filePath := "../storage/articles.json"

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalln(err)
		return ArticleData{}, err
	}

	err = json.Unmarshal(data, &wrapper)
	if err != nil {
		log.Fatalln(err)
		return ArticleData{}, err
	}

	for i := 0; i < len(wrapper.Articles); i++ {
		article := wrapper.Articles[i]
		if article.Id == articleNumber {
			return article, nil
		}
	}

	//	for _, article := range wrapper.Articles {

	//		if article.Id == articleNumber {
	//			return article, nil
	//		}
	//	}
	fmt.Println("fetching data went well")
	// handling the panic in for range
	/*
		defer func() {
			if err := recover(); err != nil {
				log.Println("panic occured:", err)
			}
		}()
	*/
	return ArticleData{}, nil
}

func GetAllArticles() ([]ArticleData, error) {

	var wrapper ArticlesWrapper
	filePath := "../storage/articles.json"

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalln(err)
		return []ArticleData{}, err
	}

	err = json.Unmarshal(data, &wrapper)
	if err != nil {
		log.Fatalln(err)
		return []ArticleData{}, err

	}
	return wrapper.Articles, nil
}
