package main

import "net/http"

// All preset responses to return.

var ResponseOK = APIResponse{http.StatusOK, ""}
var ResponseBadAPIRequest = APIResponse{http.StatusBadRequest, "bad request"}
var ResponseInvalidSession = APIResponse{http.StatusForbidden, "invalid session"}
var ResponseInvalidEndpoint = APIResponse{http.StatusBadRequest, "no such endpoint"}
var ResponseInvalidMethod = APIResponse{http.StatusBadRequest, "invalid request method"}

// resources
var ResponseUnknownResource = APIResponse{http.StatusNotFound, "unknown resource"}

// plugins
var ResponseUnknownPlugin = APIResponse{http.StatusNotFound, "unknown plugin"}