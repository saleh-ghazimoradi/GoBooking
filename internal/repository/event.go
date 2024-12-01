package repository

import (
	"context"
	"database/sql"
	"errors"
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

type eventRepo struct {
	dbWrite *sql.DB
	dbRead  *sql.DB
	tx      *sql.Tx
}

func (e *eventRepo) GetMany(ctx context.Context) ([]*service_models.Event, error) {
	return nil, nil
}

func (e *eventRepo) GetOne(ctx context.Context, id int64) (*service_models.Event, error) {
	query := `SELECT id, name, location, date, created_at, updated_at, version FROM events WHERE id = $1`

	var event service_models.Event
	err := e.dbRead.QueryRowContext(ctx, query, id).Scan(
		&event.ID,
		&event.Name,
		&event.Location,
		&event.Date,
		&event.CreatedAt,
		&event.UpdatedAt,
		&event.Version,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &event, nil
}

func (e *eventRepo) CreateOne(ctx context.Context, event *service_models.Event) error {
	query := `INSERT INTO events(name, location) VALUES ($1, $2) RETURNING id, date ,created_at, updated_at`

	err := e.dbWrite.QueryRowContext(ctx, query, event.Name, event.Location).Scan(&event.ID, &event.Date, &event.CreatedAt, &event.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (e *eventRepo) UpdateOne(ctx context.Context, event *service_models.Event) error {
	return nil
}

func (e *eventRepo) DeleteOne(ctx context.Context, id int64) error {
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
