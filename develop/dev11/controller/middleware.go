package controller

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

func (c *Controller) CheckDate(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		date := r.URL.Query().Get("date")
		_, err := time.Parse(time.DateOnly, date)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("parse date error. [YYYY.MM.DD]"))
			return
		}
		next(w, r)
	}
}

func LogMiddleware(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		end := time.Now()
		//TODO print status code
		slog.Info(fmt.Sprintf("lead_time: %d | url: %s | method: %s", end.Sub(start).Milliseconds(), r.RemoteAddr, r.Method))
	}
}

type ErrorResponse struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}

func WriteError(messageError string, status int, w http.ResponseWriter) {
	errResponse := ErrorResponse{Error: messageError, Status: status}
	data, err := json.Marshal(errResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("server error"))
	}

	w.WriteHeader(status)
	_, err = w.Write(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("server error"))
	}
}

