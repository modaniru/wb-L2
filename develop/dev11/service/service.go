package service

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

// В будущем лучше всего сделать интерфейс, но пока проектик маленький можно этип принебречь
type Service struct {
	mapa sync.Map
}

var (
	ErrEventAlreadyExists = errors.New("error: event already exists")
	ErrEventNotExists     = errors.New("error: event not exists")
)

func NewService() *Service {
	return &Service{mapa: sync.Map{}}
}

func (s *Service) AddEvent(ctx context.Context, date time.Time, event string) error {
	dateString := date.Format(time.DateOnly)
	_, ok := s.mapa.Load(dateString)
	if ok {
		return ErrEventAlreadyExists
	}
	s.mapa.Store(dateString, event)
	fmt.Println(s.mapa.Load(dateString))
	return nil
}

func (s *Service) UpdateEvent(ctx context.Context, date time.Time, event string) error {
	dateString := date.Format(time.DateOnly)
	_, ok := s.mapa.Load(dateString)
	if !ok {
		return ErrEventNotExists
	}
	s.mapa.Store(dateString, event)
	return nil
}

func (s *Service) DeleteEvent(ctx context.Context, date time.Time) error {
	dateString := date.Format(time.DateOnly)
	_, ok := s.mapa.Load(dateString)
	if !ok {
		return ErrEventNotExists
	}
	s.mapa.Delete(dateString)
	return nil
}

//TODO needs move to dto's
type EventResponse struct{
	Date string `json:"date"`
	Event string `json:"event"`
}

func (s *Service) GetDayEvents(ctx context.Context, date time.Time) (*EventResponse, error){
	dateString := date.Format(time.DateOnly)
	value, ok := s.mapa.Load(dateString)
	if !ok {
		return nil, ErrEventNotExists
	}
	return &EventResponse{Date: dateString, Event: value.(string)}, nil
}

func (s *Service) GetWeekEvents(ctx context.Context, date time.Time) ([]EventResponse, error){
	events := []EventResponse{}
	for i := 0; i < 7; i++{
		dateString := date.Format(time.DateOnly)
		value, ok := s.mapa.Load(dateString)
		if ok {
			events = append(events, EventResponse{Date: dateString, Event: value.(string)})
		}
		date = date.Add(time.Hour * 24)
	}
	if len(events) == 0 {
		return nil, ErrEventNotExists
	}
	return events, nil
}

func (s *Service) GetMonthEvents(ctx context.Context, date time.Time) ([]EventResponse, error){
	events := []EventResponse{}
	for i := 0; i < 31; i++{
		dateString := date.Format(time.DateOnly)
		value, ok := s.mapa.Load(dateString)
		if ok {
			events = append(events, EventResponse{Date: dateString, Event: value.(string)})
		}
		date = date.Add(time.Hour * 24)
	}
	if len(events) == 0 {
		return nil, ErrEventNotExists
	}
	return events, nil
}