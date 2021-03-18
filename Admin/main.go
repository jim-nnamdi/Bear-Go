package main 

// Admin structure
type Admin struct{
	ID int
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

var admintmpl = template.Must(template.ParseGlob("adminforms/*"))

func retrieveAdmins(w http.ResponseWriter, r *http.Request){
	db := dbConnection()
	results,err := db.Query("SELECT * FROM admins")
	errorCheck(err)

	adm := Admin{}
	res := []Admin{}

	for results.Next(){
		var id int 
		var name string 
		var designation string 

		err = results.Scan(&id, &name, &designation)
		errorCheck(err)

		admn.ID = id 
		adm.Name = name 
		adm.Designation = designation

		res = append(res, adm)
	}
	return res 
	defer db.Close()
}

func retrieveSingleAdmin(w http.ResponseWriter, r *http.Request){
	db := dbConnection()
	adminId := r.URL.Query().Get("id")
	result,err := db.Query("SELECT * FROM admins WHERE id=?", adminId)
	errorCheck(err)

	adm := Admin{}
	for result.Next(){
		var id int 
		var name string 
		var designation string 

		err = result.Scan(&id, &name, &designation)
		errorCheck(err)

		adm.ID = id 
		adm.Name = name 
		adm.Designation = designation
	}
	return adm 
	defer db.Close()
}

func editSingleAdminData(w http.ResponseWriter, r *http.Request){
	db := dbConnection()
	adminId := r.URL.Query().Get("id")
	result,err := db.Query("SELECT * FROM admins WHERE id=?", adminId)
	errorCheck(err)

	adm := Admin{}
	for result.Next(){
		var id int 
		var name string 
		var designation string 

		err = result.Scan(&id, &name, &designation)
		errorCheck(err)

		adm.ID = id 
		adm.Name = name 
		adm.Designation = designation
	}
	return adm 
	defer db.Close()
}