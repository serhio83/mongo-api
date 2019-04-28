package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func healthz(w http.ResponseWriter, _ *http.Request) {
	okz := struct {
		Okz string `json:"okz"`
	}{"work fine"}
	body,err := json.Marshal(okz)
	if err != nil {
		log.Println("Error marshaling json", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
