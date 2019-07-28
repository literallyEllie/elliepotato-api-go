package main

import "net/http"

// All preset responses to return.

var ResponseOK = APIResponse{http.StatusOK, ""}
var ResponseBadAPIRequest = APIResponse{http.StatusBadRequest, "bad request"}
var ResponseInvalidSession = APIResponse{http.StatusForbidden, "invalid session"}
var ResponseInvalidEndpoint = APIResponse{http.StatusBadRequest, "no such endpoint"}
var ResponseInvalidMethod = APIResponse{http.StatusBadRequest, "invalid request method"}
var ResponseRequiresAuth = APIResponse{http.StatusBadRequest, "this method requires authentication"}

// Internal error
var ResponseOopsie = APIResponse{http.StatusInternalServerError, "internal error"}

// resources
var ResponseUnknownResource = APIResponse{http.StatusNotFound, "unknown resource"}

// plugins
var ResponseUnknownPlugin = APIResponse{http.StatusNotFound, "unknown plugin"}

// service
var ResponseServiceDoesntExist = APIResponse{http.StatusNotFound, "service does not exist"}
var ResponseInvalidServiceAccessKey = APIResponse{http.StatusBadRequest, "invalid service access key"}
var ResponseIdKeyMismatch = APIResponse{http.StatusBadRequest, "mismatch of id and key"}
