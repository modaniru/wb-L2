package controller

import (
	"dev11/service"
	"encoding/json"
	"net/http"
	"time"
)

type Controller struct {
	router  http.Handler
	service *service.Service
}

func NewController(service *service.Service) *Controller {
	router := http.ServeMux{}
	c := Controller{router: &router, service: service}
	router.HandleFunc("/ping", c.Ping())
	router.HandleFunc("/event", LogMiddleware(c.CheckDate(c.EventHandler())))
	router.HandleFunc("/events_for_day", LogMiddleware(c.CheckDate(c.GetEventDay)))
	router.HandleFunc("/events_for_week", LogMiddleware(c.CheckDate(c.GetEventMonth)))
	router.HandleFunc("/events_for_month", LogMiddleware(c.CheckDate(c.GetEventWeek)))
	return &c
}

func (c *Controller) Ping() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			WriteError("endpoint supports only GET method!", 404, w)
			return
		}
		w.WriteHeader(200)
		w.Write([]byte("pong"))
	}
}

func (c *Controller) EventHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			c.AddEvent(w, r)
		case http.MethodDelete:
			c.DeleteEvent(w, r)
		case http.MethodPut:
			c.PutEvent(w, r)
		default:
			WriteError("unsupport method!", 404, w)
			return
		}

	}
}

func (c *Controller) AddEvent(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")
	t, _ := time.Parse(time.DateOnly, date)
	err := c.service.AddEvent(r.Context(), t, r.URL.Query().Get("event"))
	if err != nil {
		WriteError(err.Error(), 500, w)
		return
	}
	w.WriteHeader(200)
}

func (c *Controller) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")
	t, _ := time.Parse(time.DateOnly, date)
	err := c.service.DeleteEvent(r.Context(), t)
	if err != nil {
		WriteError(err.Error(), 500, w)
		return
	}
	w.WriteHeader(200)
}

func (c *Controller) PutEvent(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")
	t, _ := time.Parse(time.DateOnly, date)
	err := c.service.UpdateEvent(r.Context(), t, r.URL.Query().Get("event"))
	if err != nil {
		WriteError(err.Error(), 500, w)
		return
	}
	w.WriteHeader(200)
}

func (c *Controller) GetEventDay(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")
	t, _ := time.Parse(time.DateOnly, date)
	result, err := c.service.GetDayEvents(r.Context(), t)
	if err != nil {
		WriteError(err.Error(), 500, w)
		return
	}
	data, err := json.Marshal(result)
	if err != nil {
		WriteError(err.Error(), 500, w)
		return
	}

	w.WriteHeader(200)
	w.Write(data)
}

func (c *Controller) GetEventWeek(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")
	t, _ := time.Parse(time.DateOnly, date)
	result, err := c.service.GetWeekEvents(r.Context(), t)
	if err != nil {
		WriteError(err.Error(), 500, w)
		return
	}
	data, err := json.Marshal(result)
	if err != nil {
		WriteError(err.Error(), 500, w)
		return
	}

	w.WriteHeader(200)
	w.Write(data)
}

func (c *Controller) GetEventMonth(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")
	t, _ := time.Parse(time.DateOnly, date)
	result, err := c.service.GetMonthEvents(r.Context(), t)
	if err != nil {
		WriteError(err.Error(), 500, w)
		return
	}
	data, err := json.Marshal(result)
	if err != nil {
		WriteError(err.Error(), 500, w)
		return
	}

	w.WriteHeader(200)
	w.Write(data)
}



func (c *Controller) GetRouter() http.Handler {
	return c.router
}
