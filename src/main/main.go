package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	Version       = "v0.1-DEV"
	APIVersion    = "v1"
	LocalBasePath = "/" + APIVersion

	EndpointIdentify = "identify"
	EndpointPlugin   = "plugin"
	EndpointService  = "service"
)

// Struct for any response from the API server
type APIResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// API server for me to do random things and will develop into something.
func main() {
	log.Println("Loading elliepotato-api version", Version)

	router := mux.NewRouter()

	// Main handle
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		WriteAPIResponse(writer, ResponseBadAPIRequest)
	})

	// Load services.
	LoadServices()
	loadCredentials()

	// v1
	router.HandleFunc(LocalBasePath, APIv1)
	router.HandleFunc(ResourcePath, HandleResourceRequest)

	http.ListenAndServe(":80", router)

	close(stopChecker)
}

// Allows for a quick responder when responding to requests.
func WriteAPIResponse(writer http.ResponseWriter, response APIResponse) {
	formattedResponse, _ := json.Marshal(&response)
	writer.WriteHeader(response.Code)
	writer.Write(formattedResponse)
}
