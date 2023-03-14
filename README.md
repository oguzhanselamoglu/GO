# GO
There are several projects for learning purposes.

# go mod init project-name
The go mod init command creates a go.mod file to track your code's dependencies
So far, the file includes only the name of your module and the Go version your code supports. But as you add dependencies, the go.mod file will list the versions your code depends on. This keeps builds reproducible and gives you direct control over which module versions to use.

# get dependencies for code in the current directory.
go get .

# run current directory
go run .

# go get -u github.com/gorilla/mux
Package gorilla/mux implements a request router and dispatcher for matching incoming requests to their respective handler.

# go get github.com/lib/pq
A pure Go postgres driver for Go’s database/sql package.

# go get github.com/joho/godotenv
We are going to use godotenv package to read the .env file. The .env file is used to save the environment variables. Environment variables is used to keep the sensitive data safe. 
