package database

import (
	"fmt"
	"log"

	"gogenggo/config"
	"gogenggo/modules/secret"

	"github.com/jmoiron/sqlx"
)

func Prepare(query string) *sqlx.Stmt {
	if config.Configs.DB.Setting.IsMaintenance {
		log.Println("[DB - Prepare] Error when preparing query:", query, ", DB is maintenance mode.")
		return nil
	}

	if DB == nil {
		log.Fatalln("[DB - Prepare] Error when preparing query:", query, ", nil mysql value.")
		return nil
	}

	stmt, err := DB.Preparex(query)
	if err != nil {
		log.Println("[DB - Prepare] Error when preparing query:", query, ", err:", err)
		return nil
	}

	return stmt
}

// Private function
func getConnection() string {
	return fmt.Sprintf(
		config.Configs.DB.Connection.Scheme,
		secret.SecretObjects.DB.Host,
		secret.SecretObjects.DB.Port,
		secret.SecretObjects.DB.User,
		secret.SecretObjects.DB.Password,
		secret.SecretObjects.DB.DBName,
		secret.SecretObjects.DB.SSLMode,
	)
}
