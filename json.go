package main

import (
	"log"
	"net/http"
	"encoding/json"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}){
	dat, err := json.Marshal(payload)

	if err != nil {
		log.Println("Failed to marshall JSON response %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}