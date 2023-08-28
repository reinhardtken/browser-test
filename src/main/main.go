package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

/*
build
export GOARCH=amd64
export GOOS=linux
go build -o aldServer  ./server.go
/Users/bytedance/sdk/go1.19.5/bin/go  build -o webXRServer  ./main.go

set path=C:\Go1.12\bin;%path%
set GOROOT=C:\Go1.12
set GOARCH=amd64
set GOOS=linux

go build -o aldServer  ./server.go
*/

func getCurrentPath() string {
	if ex, err := os.Executable(); err == nil {
		return filepath.Dir(ex)
	}
	return "./"
}

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
	router.Static("/webxr/", `/Users/bytedance/workspace/code/self/browser-test/html/webxr-samples`)
	router.Static("/webgl/", `/Users/bytedance/workspace/code/self/browser-test/html/webglbook`)
	router.Static("/webgl2/", `/Users/bytedance/workspace/code/self/browser-test/html/WebGL2Samples`)

	//linux
	//router.Static("/static/", `/home/liuqingping/workspace/tmp/weblayer_test/html`)
	//router.Static("/js", `/home/liuqingping/workspace/tmp/weblayer_test/html/js`)
	//router.Static("/webxr/", `/home/liuqingping/workspace/tmp/weblayer_test/html/webxr`)

	//pwd
	//path := getCurrentPath()
	//fmt.Printf("webxrtest server path %s \n", path)
	//router.Static("/static/", path+`/html`)
	//router.Static("/js", path+`/html/js`)
	//router.Static("/webxr/", path+`/html/webxr`)

	router.GET("/hello",
		func(context *gin.Context) {
			context.String(http.StatusOK, "<p>hello world</p>")
		})

	fmt.Printf("webxrtest server run...")
	router.Run(":80")
}
