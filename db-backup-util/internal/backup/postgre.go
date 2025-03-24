package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func Connect() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), "postgres://myuser:mypassword@localhost:7899/mydatabase")
	if err != nil {
		return nil, err
	}
	return conn, nil
}

//func backupPostgre(conn *pgx.Conn, outPutFile string) error {

//}

func QueryUserData(conn *pgx.Conn) {
	rows, err := conn.Query(context.Background(), "select id, name from users")
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)

		if err != nil {
			log.Println(err)
		}
		fmt.Printf("User ID: %d, Name: %s\n", id, name)
	}
}
