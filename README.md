https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin

This is a simple expmple to show case creating a simple REST api using:
Gin-swagger

Create a project directory, clone a empty repo e.g git checkout github.com/nidumolu/simple-gin-swagger
cd simple-gin-swagger
git checkout github.com/nidumolu/simple-gin-swagger
go mod init github.com/nidumolu/simple-gin-swagger


// import nelow required packages
//Starting in Go 1.17, installing executables with go get is deprecated. go install may be used instead:
go install github.com/swaggo/swag/cmd/swag@latest

go get -u github.com/gin-gonic/gin
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files

swag init

go build

go run main.go

Test it out @ http://localhost:8080/swagger/index.html