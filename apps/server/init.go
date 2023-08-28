package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"gogenggo/apps/server/handlers"
	"gogenggo/apps/server/routes"
	"gogenggo/config"
	"gogenggo/internals/usecases"
)

func InitServer(usecase *usecases.UsecaseWrapper) (*Server, error) {
	srv := new(Server)

	handler, err := handlers.Init(usecase)
	if err != nil {
		log.Fatal("[Server - InitServer] Error initializing server, err:", err)
		return nil, err
	}

	route, err := routes.Init(handler)
	if err != nil {
		log.Fatal("[Server - InitServer] Error initializing routes, err:", err)
		return nil, err
	}

	srv.Handlers = handler
	srv.Routes = route
	srv.Port, _ = strconv.Atoi(os.Getenv("PORT"))
	if srv.Port == 0 {
		srv.Port = config.Configs.Main.Http.Port
	}

	srv.Http = &http.Server{
		Handler:      route.Mux,
		Addr:         fmt.Sprintf(":%d", srv.Port),
		WriteTimeout: time.Duration(config.Configs.Main.Http.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(config.Configs.Main.Http.ReadTimeout) * time.Second,
	}

	return srv, nil
}
