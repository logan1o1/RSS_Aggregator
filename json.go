package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(resp http.ResponseWriter, statusCode int, msg string) {
	if statusCode > 499 {
		log.Println("Responding with 500 error: ", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(resp, statusCode, errResponse{
		Error: msg,
	})
}

func respondWithJSON(resp http.ResponseWriter, statusCode int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal json resp %v", payload)
		resp.WriteHeader(500)
		return
	}
	resp.Header().Add("Content-Type", "application/json")
	resp.WriteHeader(statusCode)
	resp.Write(data)
}
