package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Service struct {
	ID        string `json:"id"`
	AccessKey string `json:"access_key"`
}

type ActiveService struct {
	ID       string
	LastPing time.Time
}

// Map for all possible services, indexed by ID
var services = map[string]*Service{}

// Map for all active services, indexed by id
var activeServices = map[string]*ActiveService{}

// Subscribers - Service[]subs
var subscribers = map[string][]string{}

var stopChecker chan struct{}

func LoadServices() {

	file, err := os.Open("cfg/services.cfg")
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

	// Temp map of services indexed by id
	var deseralizedServices map[string]*Service

	err = json.Unmarshal(bData, &deseralizedServices)
	if err != nil {
		LogErr(err)
	}

	for _, v := range deseralizedServices {
		services[v.ID] = v
	}

	activeChecker := time.NewTicker(5 * time.Minute)
	stopChecker = make(chan struct{})

	go func() {
		for {
			select {
			case <-activeChecker.C:
				doActiveCheck()
			case <-stopChecker:
				activeChecker.Stop()
				return
			}
		}
	}()

}

func HandleServiceEndpoint(request APIRequest) APIResponse {

	// Subscribe for death updates.
	if request.Method == "Subscribe" || request.Method == "GetStatus" {

		pair, response := AuthLoginKey(request)

		if pair == nil {
			return response
		}

		subTo := request.Payload["to"]
		if subTo == "" {
			return ResponseBadAPIRequest
		}

		// Service indexed by service key
		service, ok := services[subTo]
		if !ok {
			return ResponseServiceDoesntExist
		}

		if request.Method == "Subscribe" {

			subscribers[service.ID] = append(subscribers[service.ID], request.IP)

			// TODO unsubscribe?

			return ResponseOK
		}

		if request.Method == "GetStatus" {

			activeService, ok := activeServices[subTo]

			if !ok {
				return APIResponse{http.StatusOK, "not_active"}
			}

			activeSerialized, err := json.Marshal(activeService)
			if err != nil {
				LogErr(err)
				return ResponseOopsie
			}

			return APIResponse{http.StatusOK, string(activeSerialized)}
		}

		return response
	}

	// Get ID
	id, ok := request.Payload["id"]
	if !ok {
		return ResponseBadAPIRequest
	}

	// Check if service exists
	service, ok := services[id]
	if !ok {
		return ResponseServiceDoesntExist
	}

	// Get Access Key
	key, ok := request.Payload["access_key"]
	if !ok {
		return ResponseBadAPIRequest
	}

	// Check if access key matches id
	if service.AccessKey != key {
		return ResponseIdKeyMismatch
	}

	// From this point it is valid.

	if request.Method == "Status" {

		status := request.Payload["type"]

		if status == "alive" {
			if activeServices[id] == nil {
				activeServices[id] = &ActiveService{id, time.Now()}
			}

			activeService := activeServices[id]
			activeService.LastPing = time.Now()

			LogInfo(activeService.ID + " registered ALIVE status from " + request.IP)
		}

		if status == "exit" {
			delete(activeServices, id)
			LogInfo(id + " unregistered EXIT status from " + request.IP)
		}

		return ResponseOK
	}

	return ResponseInvalidMethod
}

func GetSubscribersTo(service string) []string {
	return subscribers[service]
}

func IsSubscribedTo(subscriber string, service string) bool {

	serviceSubs := GetSubscribersTo(service)

	for _, sub := range serviceSubs {
		if sub == subscriber {
			return true
		}
	}

	return false
}

func NotifySubscribersOf(serviceId string) {

	// How do we verify the sender that we are who we are?

}

func doActiveCheck() {

	for _, service := range activeServices {

		LogInfo(fmt.Sprintf("debug checking %s from %s", service.ID, service.LastPing))

		if time.Since(service.LastPing) > time.Minute*5 {
			LogWarn(fmt.Sprintf("No response from %s in 5 minute. Presuming dead.", service.ID))

			delete(activeServices, service.ID)
		}

	}

}
