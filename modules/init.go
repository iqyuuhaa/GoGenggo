package modules

import (
	"log"

	"gogenggo/modules/cron"
	"gogenggo/modules/database"
	"gogenggo/modules/secret"
)

func InitAllModules() error {
	if err := secret.Init(); err != nil {
		log.Fatalln("[Modules - InitAllModules] Error initializing secret data, err: ", err)
		return err
	}

	if err := database.Init(); err != nil {
		log.Fatalln("[Modules - InitAllModules] Error initializing database, err: ", err)
		return err
	}

	cron.InitCron()

	return nil
}
