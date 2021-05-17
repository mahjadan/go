package routes

import (
	"echo-demo-print-routes/handlers"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type Route struct {
	Path    string
	Method  string
	Handler echo.HandlerFunc
}

type HandlerConfig struct {
	AccountHandler  handlers.AccountHandler
	CustomerHandler handlers.CustomerHandler
	HealthHandler   handlers.HealthHandler
}

func SetupRoutes(e *echo.Echo, c HandlerConfig) {
	AccRoutes := accountRoutes(c.AccountHandler)
	cusRoutes := customerRoutes(c.CustomerHandler)
	healthRoutes := healthRoutes(c.HealthHandler)

	var routes []*Route
	routes = append(routes, AccRoutes...)
	routes = append(routes, cusRoutes...)
	routes = append(routes, healthRoutes...)
	install(e, routes)
}

func install(e *echo.Echo, routes []*Route, CORSMiddleware ...echo.MiddlewareFunc) {
	for _, route := range routes {
		if corsSkipper(route) {
			e.Add(route.Method, route.Path, route.Handler)
		} else {
			e.Add(route.Method, route.Path, route.Handler, CORSMiddleware...)
		}
	}
}

func corsSkipper(route *Route) bool {
	if strings.HasPrefix(route.Path, "/accounts") ||
		strings.HasPrefix(route.Path, "/buyers") ||
		strings.HasPrefix(route.Path, "/fast-registry") && route.Method == http.MethodPost ||
		strings.HasPrefix(route.Path, "/resource-status") ||
		strings.HasPrefix(route.Path, "/login") ||
		strings.HasPrefix(route.Path, "/health") ||
		strings.HasPrefix(route.Path, "/check") {

		return true
	}
	return false
}

func accountRoutes(controller handlers.AccountHandler) []*Route {
	return []*Route{
		{
			Path:    "/accounts",
			Method:  http.MethodGet,
			Handler: controller.List,
		},
	}
}

func customerRoutes(controller handlers.CustomerHandler) []*Route {
	return []*Route{
		{
			Path:    "/customers",
			Method:  http.MethodGet,
			Handler: controller.GetById,
		},
	}
}

func healthRoutes(controller handlers.HealthHandler) []*Route {
	return []*Route{
		{
			Path:    "/health",
			Method:  http.MethodGet,
			Handler: controller.Get,
		},
	}
}
