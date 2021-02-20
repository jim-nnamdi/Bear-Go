package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	ID          int
	Name        string
	Description string
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

func dbConn() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:root@/gotest")
	errorCheck(err)
	return db
}

var tmpl = template.Must(template.ParseGlob("forms/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	results, err := db.Query("SELECT * FROM posts")
	errorCheck(err)

	pst := Post{}
	res := []Post{}

	for results.Next() {
		var ID int
		var Name string
		var Description string

		err := results.Scan(&ID, &Name, &Description)
		errorCheck(err)
		pst.ID = ID
		pst.Name = Name
		pst.Description = Description

		res = append(res, pst)

	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}
