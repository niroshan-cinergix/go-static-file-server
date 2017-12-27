package main

import (
	"net/http"
	"os"
	"path"
)

func targetFile(request *http.Request) string {
	fileOnRequest := request.URL.Path
	targetName := path.Join("./pages", fileOnRequest)
	f, err := os.Open(targetName)
	if err != nil {
		return "./pages/404.html"
	}
	defer f.Close()
	return targetName
}

func staticFileServer() http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		http.ServeFile(response, request, targetFile(request))
	})
}

func main() {
	http.ListenAndServe(":8080", staticFileServer())
}
