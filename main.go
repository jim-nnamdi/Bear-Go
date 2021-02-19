package main 

import (
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

func main(){
	fmt.Println("connection to the Database")

	db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")

	if err != nil{
		panic (err.Error())
	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO test VALUES(2, 'Test')")

	if err != nil {
		panic (err.Error())
	}

	defer insert.Close()
}