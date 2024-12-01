package gateway

import (
	"github.com/saleh-ghazimoradi/GoBooking/logger"
	"net/http"
)

func internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	logger.Logger.Error("internal error", "method", r.Method, "path", r.URL.Path, "err", err.Error())
	writeJSONError(w, http.StatusInternalServerError, "the server encountered an error")
}

func badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	logger.Logger.Warn("bad request", "method", r.Method, "path", r.URL.Path, "err", err.Error())
	writeJSONError(w, http.StatusBadRequest, err.Error())
}

func notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	logger.Logger.Warn("not found", "method", r.Method, "path", r.URL.Path, "err", err.Error())
	writeJSONError(w, http.StatusNotFound, "not found")
}

func conflictResponse(w http.ResponseWriter, r *http.Request, err error) {
	logger.Logger.Error("conflict response", "method", "path", "error", r.Method, r.URL.Path, err.Error())
	writeJSONError(w, http.StatusConflict, err.Error())
}

func unauthorizedErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	logger.Logger.Warn("unauthorized error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJSONError(w, http.StatusUnauthorized, "unauthorized")
}

func unauthorizedBasicErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	logger.Logger.Warn("unauthorized basic error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	writeJSONError(w, http.StatusUnauthorized, "unauthorized")
}

func forbiddenResponse(w http.ResponseWriter, r *http.Request) {
	logger.Logger.Warn("forbidden", "method", r.Method, "path", r.URL.Path, "error")
	writeJSONError(w, http.StatusForbidden, "forbidden")
}

func rateLimitExceededResponse(w http.ResponseWriter, r *http.Request, retryAfter string) {
	logger.Logger.Warn("rate limit exceeded", "method", r.Method, "path", r.URL.Path)
	w.Header().Set("Retry_After", retryAfter)
	writeJSONError(w, http.StatusTooManyRequests, "rate limit exceeded, retry after: "+retryAfter)
}
