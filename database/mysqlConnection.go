package database

import (
	"DataRetriever/model"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func CreateConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/test")
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("db is connected")
		defer db.Close()

		err = db.Ping()
		fmt.Println(err)
		if err != nil {
			fmt.Println("MySQL db is not connected")
			fmt.Println(err.Error())
		}
	}
	return db
}

func InsertPost(db *sql.DB, post model.Post) {
	stmtIns, err := db.Prepare("INSERT INTO squareNum VALUES( ?, ? )") // ? = placeholder
	if err != nil {
		panic(err.Error())
	}

	_, err = stmtIns.Exec(post.Title, post.Comments, post.RetrievedAt, post.StoryURL)
	if err != nil {
		panic(err.Error())
	}

	defer stmtIns.Close()
}
