package gateway

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/saleh-ghazimoradi/GoBooking/config"
	"github.com/saleh-ghazimoradi/GoBooking/docs"
	_ "github.com/saleh-ghazimoradi/GoBooking/docs"
	"github.com/saleh-ghazimoradi/GoBooking/internal/repository"
	"github.com/saleh-ghazimoradi/GoBooking/internal/service"
	"github.com/saleh-ghazimoradi/GoBooking/utils"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

// @title GoBooking API
// @version 1.0
// @description This is a web API server.
// @BasePath /
// @schemes http https
// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
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

	swaggerHandler := SetupSwagger()
	router.Handler(http.MethodGet, "/swagger/*any", swaggerHandler)

	return router
}

func SetupSwagger() http.Handler {
	docs.SwaggerInfo.Title = "Golang Web API"
	docs.SwaggerInfo.Description = "This is a web API server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost%s", config.Appconfig.Server.Port)
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	return httpSwagger.WrapHandler
}
