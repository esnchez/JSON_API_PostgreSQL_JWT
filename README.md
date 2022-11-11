## API development in Golang

Implementing a Json API in Golang

We will use gorilla/mux pachage for creating our HTTP router. 
Add it to mod file with: go get github.com/gorilla/mux

We run a Postgres image that our API will connect to. Adjust your Postgres image tag.    
(sudo) docker run --name some-postgres -e POSTGRES_PASSWORD=1234 -p 5432:5432 -d postgres:alpine

