package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type CredentialsPair struct {
	Username    string   `json:"username"`
	LoginKey    string   `json:"login_key"`
	Permissions []string `json:"permissions"`
}

var credentials map[string]CredentialsPair

func loadCredentials() {
	file, err := os.Open("priv/credentials.cfg")

	if err != nil {
		LogErr(err)
		return
	}
	defer file.Close()

	bData, err := ioutil.ReadAll(file)
	if err != nil {
		LogErr(err)
		return
	}

	err = json.Unmarshal(bData, &credentials)
	if err != nil {
		LogErr(err)
		return
	}

	LogInfo(fmt.Sprintf("Loaded %v credential pairs", len(credentials)))
}

func AuthLoginKey(request APIRequest) (*CredentialsPair, APIResponse) {

	loginKey := request.Payload["login_key"]
	if loginKey == "" {
		return nil, ResponseRequiresAuth
	}

	id := GetCredentials(loginKey)

	if id == nil {
		return nil, ResponseRequiresAuth
	}

	return id, ResponseOK
}

func GetCredentials(loginKey string) *CredentialsPair {

	var id *CredentialsPair
	for _, v := range credentials {
		if v.LoginKey == loginKey {
			id = &v
		}
	}

	return id
}
