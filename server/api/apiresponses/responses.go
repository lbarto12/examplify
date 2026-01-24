package apiresponses

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/gommon/log"
)

func Error(w http.ResponseWriter, message string, status int) {
	http.Error(w, message, status)
}

func ErrorWithBody(w http.ResponseWriter, body any, status int) {
	res, err := json.Marshal(body)
	if err != nil {
		http.Error(w, "error marshalling response", http.StatusInternalServerError)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(res)
	if err != nil {
		log.Error(err)
	}
}

func Success(w http.ResponseWriter, response any) {
	res, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "error marshalling response", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		log.Error(err)
	}
}
