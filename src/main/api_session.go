package main

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"time"
)

// sessionId[Session]
var sessions = map[string]APISession{}

// GetSession gets the session from a sessionId string
// Returns an APISession and whether the APISession exists or not.
func GetSession(sessionId string) (APISession, bool) {
	session, ok := sessions[sessionId]
	return session, ok
}

func Authenticated(request APIRequest) bool {
	if request.SessionKey == "" {
		return false
	}

	// TODO

	return false
}

// Creates a session through a request struct.
// If the request possibly has an API key, it will be invalidated and reset.
// Returns a newly created API Session.
func createSession(apiRequest APIRequest) APISession {
	if apiRequest.SessionKey != "" {
		invalidateSession(apiRequest.SessionKey)
	}

	session := APISession{
		SessionKey: generateNewSessionKey(),
		Created:    time.Now(),
	}

	sessions[session.SessionKey] = session
	return session
}

// Invalidates a session through a sessionId string by deleting it from the session map.
func invalidateSession(sessionId string) {
	_, ok := GetSession(sessionId)

	if ok {
		delete(sessions, sessionId)
		return
	}
}

// Generates and returns a unique session key.
func generateNewSessionKey() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
