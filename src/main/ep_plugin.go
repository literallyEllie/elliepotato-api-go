package main

import "net/http"

// Contains a map of all the plugin versions I have.
var pluginVersions = map[string]string{
	"cmd-alias": "1.3.1",
	"sleepy":    "1.1-RELEASE",
}

// Handler for the plugin end point allowing to get the most up-to-date version.
func HandlePluginEndpoint(request APIRequest) APIResponse {

	if request.Method == "GetVersion" {
		toCheckPlugin := request.Payload["plugin-id"]

		currentVersion, validPlugin := pluginVersions[toCheckPlugin]

		if !validPlugin {
			return ResponseUnknownPlugin
		}

		return APIResponse{http.StatusOK, currentVersion}
	}

	return ResponseInvalidMethod
}
