package gateway

import (
	"context"
	"errors"
	"fmt"
	"github.com/saleh-ghazimoradi/GoBooking/config"
	"github.com/saleh-ghazimoradi/GoBooking/logger"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var wg sync.WaitGroup

func Server(router http.Handler) error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.Appconfig.Server.Port),
		Handler:      router,
		IdleTimeout:  config.Appconfig.Server.IdleTimeout,
		ReadTimeout:  config.Appconfig.Server.ReadTimeout,
		WriteTimeout: config.Appconfig.Server.WriteTimeout,
	}

	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit

		logger.Logger.Info("shutting down server", "signal", s.String())

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		err := srv.Shutdown(ctx)
		if err != nil {
			shutdownError <- err
		}

		logger.Logger.Info("completing background tasks", "addr", srv.Addr)

		wg.Wait()
		shutdownError <- nil
	}()

	logger.Logger.Info("starting server", "addr", srv.Addr, "env", config.Appconfig.Server.Version)

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownError
	if err != nil {
		return err
	}

	logger.Logger.Info("stopped server", "addr", srv.Addr)

	return nil
}
