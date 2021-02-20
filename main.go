package main

import (
	"database/sql"
	"log"
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

func index(w http.ResponseWriter, r *http.Request) {
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

func show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	postID := r.URL.Query().Get("ID")

	result, err := db.Query("SELECT * FROM posts WHERE id=?", postID)
	errorCheck(err)

	pst := Post{}
	for result.Next() {
		var ID int
		var Name string
		var Description string

		err := result.Scan(&ID, &Name, &Description)
		errorCheck(err)

		pst.ID = ID
		pst.Name = Name
		pst.Description = Description

		tmpl.ExecuteTemplate(w, "Show", pst)
		defer db.Close()
	}
}

func new(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	postID := r.URL.Query().Get("ID")
	result, err := db.Query("SELECT * FROM posts WHERE id=?", postID)
	errorCheck(err)

	pst := Post{}
	for result.Next() {
		var ID int
		var Name string
		var Description string

		err := result.Scan(&ID, &Name, &Description)
		errorCheck(err)

		pst.ID = ID
		pst.Name = Name
		pst.Description = Description
	}

	tmpl.ExecuteTemplate(w, "Edit", pst)
	defer db.Close()
}

func delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	postID := r.URL.Query().Get("ID")
	stmt, err := db.Prepare("DELETE * FROM posts WHERE id=?")
	errorCheck(err)

	stmt.Exec(postID)

	log.Println("Resource deleted")
	defer db.Close()
	http.Redirect(w, r, "/", 301)

}

func insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")

		insform, err := db.Prepare("INSERT INTO posts (name, description) VALUES(?,?)")
		errorCheck(err)

		insform.Exec(name, description)
		log.Println("Resource Added" + name + " " + description)
	}

	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		id := r.FormValue("uid")

		stmt, err := db.Prepare("UPDATE posts SET name=?, description=? WHERE id=?")
		errorCheck(err)

		stmt.Exec(name, description, id)
		log.Println("Resource updated !")
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func main() {
	log.Println("server started")
	http.HandleFunc("/", index)
	http.HandleFunc("/show", show)
	http.HandleFunc("/new", new)
	http.HandleFunc("/edit", edit)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", delete)
	http.ListenAndServe(":8080", nil)
}
