package main


import (

"github.com/gin-gonic/gin"
	"net/http"
)


/*
build
export GOARCH=amd64
export GOOS=linux
go build -o aldServer  ./server.go

set path=C:\Go1.12\bin;%path%
set GOROOT=C:\Go1.12
set GOARCH=amd64
set GOOS=linux

go build -o aldServer  ./server.go
*/


func main() {

	//router1################################################
	router := gin.New()
	router.Static("/static/", `D:/workspace/code/self/browser-test/html`)
	router.Static("/js", `D:/workspace/code/self/browser-test/html/js`)
	router.GET("/hello",
		func(context *gin.Context) {
			context.String(http.StatusOK, "<p>hello world</p>")
		})

	router.Run(":8800")
}

