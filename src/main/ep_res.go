package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

const ResourcePath = LocalBasePath + "/res/"

// Handler for providing file resources - separate from main API.
func HandleResourceRequest(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path[len(ResourcePath):]

	file, err := ioutil.ReadFile("public-resources/" + string(path))
	if err != nil {
		WriteAPIResponse(w, ResponseUnknownResource)
		return
	}

	var contentType string

	if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png"
	} else if strings.HasSuffix(path, ".jpg") {
		contentType = "image/jpg"
	} else if strings.HasSuffix(path, ".js") {
		contentType = "text/javascript"
	} else {
		contentType = "text/plain"
	}

	w.Header().Add("Content-Type", contentType)
	w.Write(file)
}
