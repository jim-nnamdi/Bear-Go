package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Post structure
type Post struct {
	ID          int
	Name        string
	Description string
	CreatedAt   string
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
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/gotest")
	errorCheck(err)
	return db
}

var tmpl = template.Must(template.ParseGlob("forms/*"))

func index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	results, err := db.Query("SELECT * FROM post")
	errorCheck(err)

	pst := Post{}
	res := []Post{}

	for results.Next() {
		var id int
		var name string
		var description string
		var createdAt string

		err = results.Scan(&id, &name, &description, &createdAt)
		errorCheck(err)
		pst.ID = id
		pst.Name = name
		pst.Description = description
		pst.CreatedAt = createdAt

		res = append(res, pst)

	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

func show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nID := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM Post WHERE id=?", nID)
	if err != nil {
		panic(err.Error())
	}
	emp := Post{}
	for selDB.Next() {
		var id int
		var name, description, createdAt string
		err = selDB.Scan(&id, &name, &description, &createdAt)
		if err != nil {
			panic(err.Error())
		}
		emp.ID = id
		emp.Name = name
		emp.Description = description
		emp.CreatedAt = createdAt
	}

	tmpl.ExecuteTemplate(w, "Show", emp)
	defer db.Close()
}

func new(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	postID := r.URL.Query().Get("id")
	result, err := db.Query("SELECT * FROM post WHERE id=?", postID)
	errorCheck(err)

	pst := Post{}
	for result.Next() {
		var id int
		var name, description, createdAt string

		err := result.Scan(&id, &name, &description, &createdAt)
		errorCheck(err)

		pst.ID = id
		pst.Name = name
		pst.Description = description
		pst.CreatedAt = createdAt
	}

	tmpl.ExecuteTemplate(w, "Edit", pst)
	defer db.Close()
}

func delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	postID := r.URL.Query().Get("id")
	stmt, err := db.Prepare("DELETE FROM Post WHERE id=?")
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

		insform, err := db.Prepare("INSERT INTO post (name, description, createdAt) VALUES(?,?,?)")
		errorCheck(err)

		insform.Exec(name, description, time.Now().Format("2006-01-02"))
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

		stmt, err := db.Prepare("UPDATE post SET name=?, description=? WHERE id=?")
		errorCheck(err)

		stmt.Exec(name, description, id)
		log.Println("Resource updated !")
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func main() {
	log.Println("server started at ")
	http.HandleFunc("/", index)
	http.HandleFunc("/show", show)
	http.HandleFunc("/new", new)
	http.HandleFunc("/edit", edit)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", delete)

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	http.ListenAndServe(":8080", nil)
}

//reflex -r '\.go' -s -- sh -c "go run main.go"
