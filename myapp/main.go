package main

import (
	_ "github.com/lib/pq"
	"log"
	"net/http"

	"github.com/eomolo/user_auth/myapp/controllers"
)

func main() {
	//Provide address server will be provided on
	log.Println("Server started on: http://localhost:8080")
	//Create and serve a route for the Index function
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/show", controllers.Show)
	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/update", controllers.Update)
	http.HandleFunc("/delete", controllers.Delete)
	http.ListenAndServe(":8080", nil)
}
