package main

import (
	"fmt"
	"log"

	postgres "github.com/ArminEbrahimpour/db-backup-util/internal/backup"
)

func main() {

	conn, err := postgres.Connect()
	if err != nil {
		log.Println(err)

	}
	postgres.QueryUserData(conn)
	fmt.Println("Query")
}
