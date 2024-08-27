package main

import (
	"encoding/json"
	"net/http"
)

func (app *Application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers http.Header) error {
	json, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Set the content type only if no header is provided
	if headers == nil {
		w.Header().Set("Content-Type", "application/json")
	}

	// Add the provided headers
	for key, values := range headers {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.WriteHeader(status)
	w.Write(json)

	return nil

}
