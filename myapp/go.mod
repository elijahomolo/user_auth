module github.com/eomolo/user_auth/myapp

go 1.17

require (
	github.com/eomolo/user_auth/myapp/models v0.0.0
	github.com/joho/godotenv v1.4.0
	github.com/lib/pq v1.10.4
)

require github.com/gorilla/mux v1.8.0 // indirect

replace github.com/eomolo/user_auth/myapp/models v0.0.0 => ./models
