package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type (
	// Struct for all incoming requests to the main API to be parsed into.
	APIRequest struct {
		SessionKey string            `json:"session"`
		Endpoint   string            `json:"endpoint"`
		Method     string            `json:"method"`
		Payload    map[string]string `json:"payload"`

		IP string
	}

	// Struct for all sessions - won't be created for every connection.
	APISession struct {
		Created    time.Time
		SessionKey string
	}
)

// Handler for the v1 endpoint. Directs requests to their endpoint.
func APIv1(w http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)

	var parsedReq APIRequest
	decodeErr := decoder.Decode(&parsedReq)

	if decodeErr != nil {
		LogErr("Error decoding request " + decodeErr.Error() + " (from " + req.Host + ")")
		WriteAPIResponse(w, ResponseBadAPIRequest)
		return
	}

	LogInfo("Incoming request from " + req.Host + " (Endpoint: " + parsedReq.Endpoint + ", Method: " + parsedReq.Method + ")")

	if parsedReq.Endpoint == "" || parsedReq.Method == "" {
		WriteAPIResponse(w, ResponseBadAPIRequest)
		return
	}

	parsedReq.IP = req.RemoteAddr

	var response APIResponse

	// Direct to endpoint
	switch parsedReq.Endpoint {
	case EndpointIdentify:
		response = HandleIdentifyEndpoint(parsedReq)
		break
	case EndpointPlugin:
		response = HandlePluginEndpoint(parsedReq)
		break
	case EndpointService:
		response = HandleServiceEndpoint(parsedReq)
	default:
		response = ResponseInvalidEndpoint
	}

	// rely response
	WriteAPIResponse(w, response)
}

// Allows to quickly check if a session is identified or not.
func IsIdentified(request APIRequest) bool {
	_, ok := GetSession(request.SessionKey)
	return ok
}
