package utils

import (
	"encoding/json"
	"net/http"
	"books-list/models"
)

func SendError(w http.ResponseWriter, stts int, err models.Error) {
	w.WriteHeader(stts)
	json.NewEncoder(w).Encode(err)
}

func SendSuccess(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}