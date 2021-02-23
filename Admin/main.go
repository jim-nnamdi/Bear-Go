package main

import (
	"database/sql"

	_ "github.com/go-sql-driver"
)

// Post struct
type Post struct {
	ID          int
	Name        string
	Description string
	CreatedAt   string
}

func panicErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func databasePing(db *sql.DB) {
	err := db.Ping()
	panicErrorCheck(err)
}

func dbConn() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/gotest")
	panicErrorCheck(err)
	return db
}

func checkTotalNumberOfPosts() {
	db := dbConn()

	results, err := db.Query("SELECT COUNT(id) AS countPosts FROM posts")
	panicErrorCheck(err)

	var countPosts int64
	for results.Next() {
		err := results.Scan(&countPosts)
		panicErrorCheck(err)
	}
	return countPosts, nil
}
