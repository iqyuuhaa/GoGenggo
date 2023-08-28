package routes

import (
	"log"
	"net/http"

	"gogenggo/apps/server/handlers"
	"gogenggo/apps/server/middlewares"

	"github.com/gorilla/mux"
)

func Init(handler *handlers.Handler) (*Route, error) {
	r := &Route{
		handlers: handler,
	}

	err := r.registerAllRoutes()
	if err != nil {
		log.Fatal("[Routes - Init] Error registering all routes, err:", err)
		return nil, err
	}

	return r, nil
}

func (r *Route) registerAllRoutes() error {
	router := mux.NewRouter()
	router.Use(middlewares.WrapHandler)

	v1 := router.PathPrefix("/api/v1").Subrouter()
	v1.HandleFunc("/webhook/dialogflow", r.handlers.HandleWebhookDialogflow).Methods(http.MethodPost, http.MethodOptions)
	v1.HandleFunc("/test-dialogflow", r.handlers.HandleTestDialogflow).Methods(http.MethodPost, http.MethodOptions)

	// v1.HandleFunc("/live/chat-bot", r.handlers.HandleLiveChatBot).Methods(http.MethodPost, http.MethodOptions)

	r.Mux = router

	return nil
}
