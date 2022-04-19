package routes

import (
	"github.com/eomolo/user_auth/myapp/controllers"
	//"net/http"

	"github.com/gorilla/mux"
)

func Handlers() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	//Create and serve a route for the Index function

	r.HandleFunc("/", controllers.Index)
	r.HandleFunc("/login", controllers.Login)
	r.HandleFunc("/verify", controllers.VerifyUser)
	r.HandleFunc("/show", controllers.Show)
	r.HandleFunc("/edit", controllers.Edit)
	r.HandleFunc("/new", controllers.New)
	r.HandleFunc("/insert", controllers.Insert)
	r.HandleFunc("/update", controllers.Update)
	r.HandleFunc("/delete", controllers.Delete)

	return r
}
