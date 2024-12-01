package gateway

import (
	"github.com/saleh-ghazimoradi/GoBooking/config"
	"net/http"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":  "ok",
		"env":     config.Appconfig.Server.Port,
		"version": config.Appconfig.Server.Version,
	}
	if err := jsonResponse(w, http.StatusOK, data); err != nil {
		internalServerError(w, r, err)
	}
}
