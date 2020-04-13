package database

import (
	"GoRedditScrapper/model"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func CreateConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:reddit@tcp(localhost:3306)/redditDB")
	if err != nil {
		fmt.Println(err.Error())
		defer db.Close()

		err = db.Ping()
		fmt.Println(err)
		if err != nil {
			fmt.Println("MySQL db is not connected")
			fmt.Println(err.Error())
		}
	}
	fmt.Println("db is connected")
	return db
}

func InsertPost(db *sql.DB, post model.Post) {
	stmtIns, err := db.Prepare("INSERT INTO POSTS_TABLE(title,subreddit) VALUES( ?, ?)") // ? = placeholder
	if err != nil {
		panic(err.Error())
	}

	_, err = stmtIns.Exec(post.Title, post.Subreddit)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer stmtIns.Close()
}
