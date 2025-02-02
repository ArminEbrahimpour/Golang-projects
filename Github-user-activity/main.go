package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"encoding/json"	

	"github.com/joho/godotenv"
)

var apiKey string

func parseResposne(res []byte) error {

	var events []map[string]interface{}

	err := json.Unmarshal(res, &events)

	if err != nil {
		panic(err)
	}

	for _, event := range events{


		fmt.Printf("Event ID: %v\n", event["id"])
		fmt.Printf("Type: %v\n", event["type"])
		fmt.Printf("Public: %v\n", event["public"])
		fmt.Printf("Created At: %v\n", event["created_at"])

		// Access nested fields (actor, repo, payload)
		actor := event["actor"].(map[string]interface{})
		fmt.Printf("Actor: %v (%v)\n", actor["login"], actor["id"])

		repo := event["repo"].(map[string]interface{})
		fmt.Printf("Repository: %v (%v)\n", repo["name"], repo["id"])

		payload := event["payload"].(map[string]interface{})
		switch event["type"] {
		case "WatchEvent":
			fmt.Printf("Action: %v\n", payload["action"])
		case "PushEvent":
			fmt.Printf("Push ID: %v\n", payload["push_id"])
			fmt.Printf("Branch: %v\n", payload["ref"])
			fmt.Printf("Head Commit: %v\n", payload["head"])
			fmt.Println("Commits:")
			commits := payload["commits"].([]interface{})
			for _, commit := range commits {
				commitMap := commit.(map[string]interface{})
				fmt.Printf("  SHA: %v\n", commitMap["sha"])
				author := commitMap["author"].(map[string]interface{})
				fmt.Printf("  Author: %v <%v>\n", author["name"], author["email"])
				fmt.Printf("  Message: %v\n", commitMap["message"])
				fmt.Printf("  URL: %v\n", commitMap["url"])
			}
			


		}


	}
	return nil
}

func getEvent(username string) ([]byte, error) {

	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	var null []byte

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		fmt.Printf("Error ! : %s", err)
		return null, err

	}

	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", apiKey)
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")

	// creating a client object to send the request
	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		fmt.Println(err)
		return null, nil
	}
	// closing the Body at the end
	defer res.Body.Close()

	// Getting the Body of the http response
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
		return null, err
	}

	return body, nil

}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
	}
	apiKey = os.Getenv("APIKEY")

	username := os.Args[1]

	test, _ := getEvent(username)
	parseResposne(test)
}
