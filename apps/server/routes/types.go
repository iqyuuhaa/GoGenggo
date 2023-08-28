package routes

import (
	"gogenggo/apps/server/handlers"

	"github.com/gorilla/mux"
)

type Route struct {
	Mux      *mux.Router
	handlers *handlers.Handler
}
