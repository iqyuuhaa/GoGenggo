package server

import (
	"net/http"

	"gogenggo/apps/server/handlers"
	"gogenggo/apps/server/routes"
)

type Server struct {
	Handlers *handlers.Handler
	Routes   *routes.Route
	Http     *http.Server
	Port     int
}
