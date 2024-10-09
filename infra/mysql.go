package infra

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:secret_password@tcp(127.0.0.1:3306)/udemy_slack_app")
	if err != nil {
		log.Fatal(err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("ping err")
		log.Fatal(err)

		return nil
	}

	return db
}
