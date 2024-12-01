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
	UpdateOne(ctx context.Context, event *service_models.Event) error
	DeleteOne(ctx context.Context, id int64) error
	GetWithTXT(tx *sql.Tx) Event
}

type eventService struct {
	eventRepo repository.Event
}

func (e *eventService) GetMany(ctx context.Context) ([]*service_models.Event, error) {
	return e.eventRepo.GetMany(ctx)
}

func (e *eventService) GetOne(ctx context.Context, id int64) (*service_models.Event, error) {
	return e.eventRepo.GetOne(ctx, id)
}

func (e *eventService) CreateOne(ctx context.Context, event *service_models.Event) error {
	return e.eventRepo.CreateOne(ctx, event)
}

func (e *eventService) UpdateOne(ctx context.Context, event *service_models.Event) error {
	return e.eventRepo.UpdateOne(ctx, event)
}

func (e *eventService) DeleteOne(ctx context.Context, id int64) error {
	return e.eventRepo.DeleteOne(ctx, id)
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
