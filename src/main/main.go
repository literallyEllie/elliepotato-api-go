package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
 	Version = "v0.1"
	APIVersion = "v1"
	LocalBasePath = "/api/" + APIVersion

 	EndpointIdentify = "identify"
	EndpointPlugin = "plugin"
)

// Struct for any response from the API server
type APIResponse struct {
	Code 	int		`json:"code"`
	Message string	`json:"message"`
}

// API server for me to do random things and will develop into something.
func main() {
	log.Println("Loading elliepotato-api version", Version)

	// Main handle
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		WriteAPIResponse(writer, ResponseBadAPIRequest)
	})
	// v1
	http.HandleFunc(LocalBasePath, APIv1)
	http.HandleFunc(ResourcePath, HandleResourceRequest)

	http.ListenAndServe(":8080", http.DefaultServeMux)
}

// Allows for a quick responder when responding to requests.
func WriteAPIResponse(writer http.ResponseWriter, response APIResponse) {
	formattedResponse, _ := json.Marshal(&response)
	writer.WriteHeader(response.Code)
	writer.Write([]byte(formattedResponse))
}