package repository

import (
	"context"
	"database/sql"
	"github.com/saleh-ghazimoradi/GoBooking/internal/service/service_models"
)

type Event interface {
	GetMany(ctx context.Context) ([]*service_models.Event, error)
	GetOne(ctx context.Context, id int64) (*service_models.Event, error)
	CreateOne(ctx context.Context, event *service_models.Event) error
	GetWithTXT(tx *sql.Tx) Event
}

type eventRepo struct {
	dbWrite *sql.DB
	dbRead  *sql.DB
	tx      *sql.Tx
}

func (e *eventRepo) GetMany(ctx context.Context) ([]*service_models.Event, error) {
	return nil, nil
}

func (e *eventRepo) GetOne(ctx context.Context, id int64) (*service_models.Event, error) {
	return nil, nil
}

func (e *eventRepo) CreateOne(ctx context.Context, event *service_models.Event) error {
	return nil
}

func (e *eventRepo) GetWithTXT(tx *sql.Tx) Event {
	return &eventRepo{
		dbWrite: e.dbWrite,
		dbRead:  e.dbRead,
		tx:      tx,
	}
}

func NewEventRepository(dbWrite *sql.DB, dbRead *sql.DB) Event {
	return &eventRepo{
		dbWrite: dbWrite,
		dbRead:  dbRead,
	}
}
