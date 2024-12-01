package gateway

import (
	"context"
	"errors"
	"github.com/saleh-ghazimoradi/GoBooking/internal/repository"
	"github.com/saleh-ghazimoradi/GoBooking/internal/service"
	"github.com/saleh-ghazimoradi/GoBooking/internal/service/service_models"
	"net/http"
	"time"
)

type event struct {
	eventService service.Event
}

func (e *event) getManyEvents(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	evs, err := e.eventService.GetMany(ctx)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = jsonResponse(w, http.StatusOK, evs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (e *event) getOneEvent(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id, err := readIDParam(r)
	if err != nil {
		badRequestResponse(w, r, err)
		return
	}

	ev, err := e.eventService.GetOne(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrRecordNotFound):
			notFoundResponse(w, r, err)
			return
		default:
			internalServerError(w, r, err)
			return
		}
	}

	if err = jsonResponse(w, http.StatusOK, ev); err != nil {
		internalServerError(w, r, err)
		return
	}
}

func (e *event) createEvent(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var eventPayload service_models.EventPayload
	if err := readJSON(w, r, &eventPayload); err != nil {
		badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(eventPayload); err != nil {
		badRequestResponse(w, r, err)
		return
	}

	ev := &service_models.Event{
		Name:     eventPayload.Name,
		Location: eventPayload.Location,
	}

	if err := e.eventService.CreateOne(ctx, ev); err != nil {
		internalServerError(w, r, err)
		return
	}

	if err := jsonResponse(w, http.StatusCreated, ev); err != nil {
		internalServerError(w, r, err)
		return
	}

}

func (e *event) updateEvent(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id, err := readIDParam(r)
	if err != nil {
		badRequestResponse(w, r, err)
		return
	}

	var eventPayload service_models.UpdateEventPayload

	if err = readJSON(w, r, &eventPayload); err != nil {
		badRequestResponse(w, r, err)
		return
	}

	if err = Validate.Struct(eventPayload); err != nil {
		badRequestResponse(w, r, err)
		return
	}

	ev, err := e.eventService.GetOne(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrRecordNotFound):
			notFoundResponse(w, r, err)
		default:
			internalServerError(w, r, err)
		}
		return
	}

	if eventPayload.Name != nil {
		ev.Name = *eventPayload.Name
	}

	if eventPayload.Location != nil {
		ev.Location = *eventPayload.Location
	}

	if err = e.eventService.UpdateOne(ctx, ev); err != nil {
		switch {
		case errors.Is(err, repository.ErrEditConflict):
			conflictResponse(w, r, err)
		default:
			internalServerError(w, r, err)
		}
		return
	}

	if err = jsonResponse(w, http.StatusOK, ev); err != nil {
		internalServerError(w, r, err)
	}
}

func (e *event) deleteEvent(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	id, err := readIDParam(r)
	if err != nil {
		badRequestResponse(w, r, err)
		return
	}

	if err = e.eventService.DeleteOne(ctx, id); err != nil {
		internalServerError(w, r, err)
		return
	}

	if err = jsonResponse(w, http.StatusOK, nil); err != nil {
		internalServerError(w, r, err)
		return
	}
}

func NewEventHandler(eventService service.Event) *event {
	return &event{
		eventService: eventService,
	}
}
