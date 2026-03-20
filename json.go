package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// APIs in json format

func respondWithError(w http.ResponseWriter, code int, msg string){
	if code > 499{ // 500 level error code is from our side, 400 level is on the Client side
		log.Println("Responding with 5XX error:", msg)
	}
	type errResponse struct{
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errResponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}){
	data, err := json.Marshal(payload)
	if(err != nil){
		log.Printf("Failed to marshal JSON response: %v\n", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}