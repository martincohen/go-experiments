package main

import (
	"context"
	"database/sql"
	"sqlboil-test/db/models"

	_ "github.com/lib/pq"
)

func testDB() error {
	db, err := sql.Open("postgres", "dbname=mutiny_go user=root password=root sslmode=disable")
	if err != nil {
		return err
	}
	defer db.Close()

	topics, err := models.Topics().All(context.Background(), db)
	if err != nil {
		return err
	}

	for _, topic := range topics {
		println(topic.Title)
	}

	return nil
}

func main() {
	err := testDB()
	if err != nil {
		panic(err)
	}
}
