package utils

import (
	"encoding/json"
	"net/http"
	"time"
)

type ApiFunc func(http.ResponseWriter, *http.Request) error

func WriteJson(w http.ResponseWriter, st int, v any) error {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(st)
	return json.NewEncoder(w).Encode(v)
}

// Trying to look for an alogrithm that will help me in creating a unique account number when a
//	a new user is created and will be unique i.e no others users will be using that account number

func CreateAcctNumber() (string, error) {
	defaultTime := time.Now().Format("12435678950405")
	return defaultTime[4:14], nil
}
