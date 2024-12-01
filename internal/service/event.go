package service

import (
	"context"
	"database/sql"
	"github.com/saleh-ghazimoradi/GoBooking/internal/repository"
	"github.com/saleh-ghazimoradi/GoBooking/internal/service/service_models"
)

type Event interface {
	GetMany(ctx context.Context) ([]*service_models.Event, error)
	GetOne(ctx context.Context, id int64) (*service_models.Event, error)
	CreateOne(ctx context.Context, event *service_models.Event) error
	GetWithTXT(tx *sql.Tx) Event
}

type eventService struct {
	eventRepo repository.Event
}

func (e *eventService) GetMany(ctx context.Context) ([]*service_models.Event, error) {
	return nil, nil
}

func (e *eventService) GetOne(ctx context.Context, id int64) (*service_models.Event, error) {
	return nil, nil
}

func (e *eventService) CreateOne(ctx context.Context, event *service_models.Event) error {
	return nil
}

func (e *eventService) GetWithTXT(tx *sql.Tx) Event {
	return &eventService{
		eventRepo: e.eventRepo.GetWithTXT(tx),
	}
}

func NewEventService(eventRepo repository.Event) Event {
	return &eventService{
		eventRepo: eventRepo,
	}
}
