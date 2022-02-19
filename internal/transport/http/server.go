package http

import (
	"Tazhben/auth/config"
	"context"
	"errors"
	"github.com/labstack/echo"
	"os"
)

func StartServer(ctx context.Context, errCh chan error, cfg *config.Config) {
	app := echo.New()

	pathPrefix := os.Getenv("CONFIG_PATH")

	if len(pathPrefix) == 0 {
		errCh <- errors.New("CONFIG_PATH is empty")
	}

	errCh <- app.Start(":8090")
}

func initConfig() error {
	pathPrefix := os.Getenv("CONFIG_PATH")

	if len(pathPrefix) == 0 {
		return errors.New("CONFIG_PATH is empty")
	}
}
