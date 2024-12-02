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

// @Summary Get Many Events
// @Description Retrieves a list of events with pagination, sorting, and filtering.
// @Tags events
// @Accept json
// @Produce json
// @Param limit query int false "Number of events to retrieve (default: 20)"
// @Param offset query int false "Offset for pagination (default: 0)"
// @Param sort query string false "Sort order, 'asc' or 'desc' (default: desc)"
// @Param search query string false "Search term for filtering events"
// @Param since query string false "Start date filter in RFC3339 format"
// @Param until query string false "End date filter in RFC3339 format"
// @Success 200 {array} service_models.Event "List of events"
// @Failure 400 {object} map[string]string "Validation error"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /v1/events [get]
func (e *event) getManyEvents(w http.ResponseWriter, r *http.Request) {
	p := service_models.PaginationFeedQuery{
		Limit:  20,
		Offset: 0,
		Sort:   "desc",
	}

	fq, err := p.Parse(r)
	if err != nil {
		badRequestResponse(w, r, err)
		return
	}

	if err = Validate.Struct(fq); err != nil {
		badRequestResponse(w, r, err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	evs, err := e.eventService.GetMany(ctx, fq)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = jsonResponse(w, http.StatusOK, evs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// @Summary Get One Event
// @Description Retrieves details of a specific event by ID.
// @Tags events
// @Accept json
// @Produce json
// @Param id path int true "Event ID"
// @Success 200 {object} service_models.Event "Event details"
// @Failure 404 {object} map[string]string "Event not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /v1/events/{id} [get]
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

// @Summary Create an Event
// @Description Creates a new event with the given details.
// @Tags events
// @Accept json
// @Produce json
// @Param event body service_models.EventPayload true "Event payload"
// @Success 201 {object} service_models.Event "Created event details"
// @Failure 400 {object} map[string]string "Validation error"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /v1/events [post]
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

// @Summary Update an Event
// @Description Updates an existing event by ID.
// @Tags events
// @Accept json
// @Produce json
// @Param id path int true "Event ID"
// @Param event body service_models.UpdateEventPayload true "Updated event payload"
// @Success 200 {object} service_models.Event "Updated event details"
// @Failure 400 {object} map[string]string "Validation error"
// @Failure 404 {object} map[string]string "Event not found"
// @Failure 409 {object} map[string]string "Edit conflict"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /v1/events/{id} [put]
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

// @Summary Delete an Event
// @Description Deletes an existing event by ID.
// @Tags events
// @Accept json
// @Produce json
// @Param id path int true "Event ID"
// @Success 204 "No Content"
// @Failure 404 {object} map[string]string "Event not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /v1/events/{id} [delete]
func (e *event) deleteEvent(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	id, err := readIDParam(r)
	if err != nil {
		badRequestResponse(w, r, err)
		return
	}

	if err = e.eventService.DeleteOne(ctx, id); err != nil {
		switch {
		case errors.Is(err, repository.ErrRecordNotFound):
			notFoundResponse(w, r, err)
		default:
			internalServerError(w, r, err)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func NewEventHandler(eventService service.Event) *event {
	return &event{
		eventService: eventService,
	}
}
