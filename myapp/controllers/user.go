package controllers

import (
	"github.com/eomolo/user_auth/myapp/models"
	"github.com/eomolo/user_auth/myapp/utils"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"text/template"
)

//define a template
var tmpl = template.Must(template.ParseGlob("forms/*"))

//define an Index function that includes the http write and read parameters
func Index(w http.ResponseWriter, r *http.Request) {
	//connect to the database
	db := utils.DbConn()
	// Query the database and return all rows from the user table
	rows, err := db.Query(`SELECT "user_id", "username", "city", "email" FROM public."users"`)
	//Handle any errors
	CheckError(err)
	//Define and populate a User struct from the returned data from the query
	usr := models.User{}
	//The list of Users that will be passed to the html template
	res := []models.User{}
	//Loop through each row and populate a User
	for rows.Next() {
		var id int
		var username, city, email string
		err = rows.Scan(&id, &username, &city, &email)
		CheckError(err)
		usr.Id = id
		usr.Email = email
		usr.Username = username
		usr.City = city
		res = append(res, usr)
	}
	//write to the Index template that will range through the User struct displaying a field for the returned data
	tmpl.ExecuteTemplate(w, "Index", res)
	//close the database connection
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	//Execute the template New which will pass the input to /insert
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	nId := r.URL.Query().Get("id")
	usr := getUser(nId)
	tmpl.ExecuteTemplate(w, "Edit", usr)
	//defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
	nId := r.URL.Query().Get("id")
	usr := getUser(nId)
	tmpl.ExecuteTemplate(w, "Show", usr)
}

// this needs to return a user or type User
func getUser(id string) models.User {
	//connect to database
	db := utils.DbConn()
	//run a query against the db filtering the user_id table using the passed id
	rows, err := db.Query(`SELECT * FROM public."users" WHERE "user_id"=$1`, id)
	//handle error
	CheckError(err)
	//construct a User
	usr := models.User{}
	for rows.Next() {
		var id int
		var username, city, email, password string
		err = rows.Scan(&id, &username, &city, &password, &email)
		CheckError(err)
		usr.Id = id
		usr.Username = username
		usr.Email = email
		usr.City = city
	}
	defer db.Close()
	return usr
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := utils.DbConn()
	//If it's a post request, assign a variable to the value returned in each field of the New page.
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")
		city := r.FormValue("city")
		email := r.FormValue("email")
		//prepare a query to insert the data into the database
		insForm, err := db.Prepare(`INSERT INTO public.users(username,password, city, email) VALUES ($1,$2, $3, $4)`)
		//check for  and handle any errors
		CheckError(err)
		//execute the query using the form data
		insForm.Exec(username, password, city, email)
		//print out added data in terminal
		log.Println("INSERT: Username: " + username + " | City: " + city + " | Email: " + email)
	}
	defer db.Close()
	//redirect to the index page
	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := utils.DbConn()
	if r.Method == "POST" {
		username := r.FormValue("username")
		city := r.FormValue("city")
		email := r.FormValue("email")
		password := r.FormValue("password")
		id := r.FormValue("uid")
		insForm, err := db.Prepare(`UPDATE public."users" SET username=$1, city=$2, email=$3, password=$4 WHERE "user_id"=$5`)
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(username, city, email, password, id)
		log.Println("UPDATE: Username: " + username)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := utils.DbConn()
	usr := r.URL.Query().Get("id")
	delForm, err := db.Prepare(`DELETE FROM public."users" WHERE user_id=$1`)
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(usr)
	log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
