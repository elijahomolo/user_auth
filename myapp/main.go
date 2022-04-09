package main

import (
	"github.com/eomolo/user_auth/myapp/routes"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	// Handle routes
	http.Handle("/", routes.Handlers())

	//Provide address server will be provided on
	log.Println("Server started on: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
