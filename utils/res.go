package utils

import (
	"encoding/json"
	"net/http"
)

func ResJson(w http.ResponseWriter, status int, p interface{}) {
	res, err := json.Marshal(p)
	if err != nil {
		http.Error(w, "Error Export", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(res))
}
