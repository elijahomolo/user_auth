module github.com/eomolo/user_auth/myapp

go 1.17

require (
	github.com/joho/godotenv v1.4.0
	github.com/lib/pq v1.10.4
	github.com/eomolo/user_auth/myapp/models v0.0.0
)

replace github.com/eomolo/user_auth/myapp/models v0.0.0 => ./models