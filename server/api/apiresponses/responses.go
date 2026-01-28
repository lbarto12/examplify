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
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if _, err := w.Write(res); err != nil {
		log.Error(err)
	}
}

func Success(w http.ResponseWriter, response any) {
	res, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "error marshalling response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(res); err != nil {
		log.Error(err)
	}
}

// BadRequest logs the error and returns a 400 Bad Request response
func BadRequest(w http.ResponseWriter, message string, err error) {
	if err != nil {
		log.Error("Bad Request: ", message, " - ", err)
	}
	Error(w, message, http.StatusBadRequest)
}

// InternalError logs the error and returns a 500 Internal Server Error response
func InternalError(w http.ResponseWriter, message string, err error) {
	if err != nil {
		log.Error("Internal Error: ", message, " - ", err)
	}
	Error(w, message, http.StatusInternalServerError)
}

// Unauthorized logs the error and returns a 401 Unauthorized response
func Unauthorized(w http.ResponseWriter, message string, err error) {
	if err != nil {
		log.Error("Unauthorized: ", message, " - ", err)
	}
	Error(w, message, http.StatusUnauthorized)
}
