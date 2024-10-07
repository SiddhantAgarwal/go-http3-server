package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{
		"pong": time.Now().Format(time.RFC3339),
	}

	bytes, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
