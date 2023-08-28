package apps

import (
	"fmt"
	"log"
	"net/http"

	"gogenggo/apps/server"
	"gogenggo/config"
	"gogenggo/internals"
	"gogenggo/modules"
	"gogenggo/modules/cache"
	longPolling "gogenggo/modules/long_polling"
)

func RunApps() {
	config.LoadConfig()

	runServers()
}

func runServers() {
	fmt.Println(">>> Initializing all modules...")
	if err := modules.InitAllModules(); err != nil {
		log.Fatalln("[Apps - RunApps] Error initializing all modules, err:", err)
		return
	}
	fmt.Println(">>> Finish initializing all modules...")

	fmt.Println(">>> Initializing all internals...")
	internal, err := internals.InitAllInternals()
	if err != nil {
		log.Fatalln("[Apps - RunApps] Error initializing all internals, err:", err)
		return
	}
	fmt.Println(">>> Finish initializing all internals...")

	fmt.Println(">>> Initializing & populate server cache...")
	if err := cache.InitServerCache(internal.Platform); err != nil {
		log.Fatalln("[Apps - RunApps] Error initializing server, err:", err)
		return
	}
	fmt.Println(">>> Finish initializing & populate server cache...")

	if !config.Configs.Main.System.IsActiveWebhook {
		fmt.Println(">>> Initializing long polling method...")
		if err := longPolling.Init(internal.Pkg, internal.Platform); err != nil {
			log.Fatalln("[Apps - RunApps] Error initializing long polling method, err:", err)
			return
		}
		fmt.Println(">>> Finish long polling method...")
	}

	fmt.Println(">>> Initializing server...")
	srv, err := server.InitServer(internal.Usecase)
	if err != nil {
		log.Fatalln("[Apps - RunApps] Error initializing server, err:", err)
		return
	}
	fmt.Println(">>> Finish initializing server...")

	fmt.Printf(">>> Running server with port :%d\n", srv.Port)
	if err = srv.Http.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalln("[Apps - RunApps] Error listening and serve server, err:", err)
	}
}
