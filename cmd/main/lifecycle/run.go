package lifecycle

import (
	"context"
	"errors"
	"github.com/lucaiatropulus/social/cmd/main/application"
	"github.com/lucaiatropulus/social/docs"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(mux http.Handler, app *application.Application) error {

	docs.SwaggerInfo.Version = "0.0.1"
	docs.SwaggerInfo.Host = app.Config.APP.ApiURL
	docs.SwaggerInfo.BasePath = "/api/v1"

	srv := &http.Server{
		Addr:         app.Config.APP.Address,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	shutdown := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		s := <-quit

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		defer cancel()

		app.Logger.Infow("signal caught", "signal", s.String())

		shutdown <- srv.Shutdown(ctx)
	}()

	app.Logger.Infow("Server has started at", "addr", app.Config.APP.Address, "env", app.Config.APP.Environment)

	err := srv.ListenAndServe()

	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdown

	if err != nil {
		return err
	}

	app.Logger.Infow("Server has stoped", "addr", app.Config.APP.Address, "env", app.Config.APP.Environment)
	return nil
}
