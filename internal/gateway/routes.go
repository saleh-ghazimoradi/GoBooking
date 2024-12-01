package gateway

import (
	"github.com/julienschmidt/httprouter"
	"github.com/saleh-ghazimoradi/GoBooking/internal/repository"
	"github.com/saleh-ghazimoradi/GoBooking/internal/service"
	"github.com/saleh-ghazimoradi/GoBooking/utils"
	"log"
	"net/http"
)

func registerRoutes() *httprouter.Router {
	db, err := utils.PostConnection()
	if err != nil {
		log.Fatal(err)
	}

	eventRepository := repository.NewEventRepository(db, db)
	eventService := service.NewEventService(eventRepository)
	eventHandler := NewEventHandler(eventService)

	router := httprouter.New()
	router.NotFound = http.HandlerFunc(notFoundRouter)
	router.MethodNotAllowed = http.HandlerFunc(methodNotAllowedResponse)
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", healthCheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/events/:id", eventHandler.getOneEvent)
	router.HandlerFunc(http.MethodGet, "/v1/events", eventHandler.getManyEvents)
	router.HandlerFunc(http.MethodPost, "/v1/events", eventHandler.createEvent)
	router.HandlerFunc(http.MethodDelete, "/v1/events/:id", eventHandler.deleteEvent)
	router.HandlerFunc(http.MethodPut, "/v1/events/:id", eventHandler.updateEvent)

	return router
}
