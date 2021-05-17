package main

import (
	"echo-demo-print-routes/banner"
	"echo-demo-print-routes/handlers"
	"echo-demo-print-routes/routes"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func main() {
	e := echo.New()

	//here you can construct your repository, service and create handler with dependency injection and pass them
	// as a config struct to setupRoutes.
	handlerConfig := routes.HandlerConfig{
		handlers.NewAccountHandler(),
		handlers.NewCustomerHandler(),
		handlers.NewHealthHandler(),
	}

	routes.SetupRoutes(e, handlerConfig)

	server := &http.Server{
		Addr:              ":8089",
		ReadTimeout:       2 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		WriteTimeout:      2 * time.Second,
		IdleTimeout:       2 * time.Second,
	}

	banner.ShowBanner(e, server.Addr)
	e.Logger.Fatal(e.StartServer(server))
}
