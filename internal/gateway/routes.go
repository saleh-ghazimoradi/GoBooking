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
	defer db.Close()

	eventRepository := repository.NewEventRepository(db, db)
	eventService := service.NewEventService(eventRepository)
	eventHandler := NewEventHandler(eventService)

	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/", eventHandler.getOneHandler)
	router.HandlerFunc(http.MethodGet, "/v1/", eventHandler.getManyHandler)
	router.HandlerFunc(http.MethodPost, "/v1/", eventHandler.createEventHandler)

	return router
}
