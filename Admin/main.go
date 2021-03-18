package main 

// Admin structure
type Admin struct{
	ID float64
	Name string 
	Designation string
}

func errorCheck(err error){
	if err != nil {
		panic(err.Error())
	}
}

func dbPing(db *sql.DB){
	db, err := db.Ping()
	errorCheck(err)
}

func dbConnection()(db *sql.DB){
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/gotest")
	errorCheck(err)
	return db
}