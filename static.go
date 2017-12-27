package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

func notFound(response http.ResponseWriter, request *http.Request) {
	body, _ := ioutil.ReadFile("./pages/404.html")
	fmt.Fprintf(response, string(body))
}

// func requestHandler(response http.ResponseWriter, request *http.Request) http.Handler {
// }

func staticFileServer(fileSystem http.FileSystem) http.Handler {
	fileServer := http.FileServer(fileSystem)
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		_, err := fileSystem.Open(path.Clean(request.URL.Path))
		if os.IsNotExist(err) {
			notFound(response, request)
		} else {
			fileServer.ServeHTTP(response, request)
		}
	})
}

func main() {
	// http.Handle("/", http.FileServer(http.Dir("./pages")))
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	panic(err)
	// }
	http.ListenAndServe(":8080", staticFileServer(http.Dir("./pages")))
}
