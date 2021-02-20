package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	ID          int
	Name        string
	Description string
}

func main() {
	db, err := sql.Open("mysql", "root:root@/gotest")
	errorCheck(err)

	defer db.Close()
	pingDB(db)

	// insert into DB

	stmt, err := db.Prepare("INSERT INTO posts(id, Name) VALUES(?,?)")
	errorCheck(err)
	result, err := stmt.Exec(5, "cool course")
	errorCheck(err)
	id, err := result.LastInsertId()
	errorCheck(err)
	fmt.Println("insert Data", id)

	// update DB

	stmt2, err := db.Prepare("UPDATE posts SET Name=? WHERE id=?")
	errorCheck(err)
	result2, err := stmt2.Exec("Goddamn", 5)
	errorCheck(err)
	id2, err := result2.RowsAffected()
	errorCheck(err)
	fmt.Println(id2)

	// return values

	stmt3, err := db.Query("SELECT * FROM posts")
	errorCheck(err)
	var post = Post{}
	for stmt3.Next() {
		err := stmt3.Scan(&post.ID, &post.Description, &post.Name)
		errorCheck(err)
		fmt.Println(post)
	}
}

func errorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func pingDB(db *sql.DB) {
	err := db.Ping()
	errorCheck(err)
}
