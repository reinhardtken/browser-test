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
	//windows
	//router.Static("/static/", `D:/workspace/code/self/browser-test/html`)
	//router.Static("/js", `D:/workspace/code/self/browser-test/html/js`)
	//linux
	//router.Static("/static/", `/home/liuqingping/workspace/code/self/github/browser-test/html`)
	//router.Static("/js", `/home/liuqingping/workspace/code/self/github/browser-test/html/js`)
	//mac
	router.Static("/static/", `/Users/bytedance/workspace/code/self/browser-test/html`)
	router.Static("/js", `/Users/bytedance/workspace/code/self/browser-test/html/js`)
	router.GET("/hello",
		func(context *gin.Context) {
			context.String(http.StatusOK, "<p>hello world</p>")
		})

	router.Run(":8800")
}

