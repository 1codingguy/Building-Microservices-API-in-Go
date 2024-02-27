module github.com/1codingguy/go-microservice-api/banking

go 1.21.5

replace github.com/1codingguy/go-microservice-api/banking/app v0.0.0 => ./app

require github.com/1codingguy/go-microservice-api/banking/app v0.0.0

require (
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
)
