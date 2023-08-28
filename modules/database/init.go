package database

import (
	"context"
	"log"
	"time"

	"gogenggo/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func Init() error {
	if DB != nil {
		return nil
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Duration(config.Configs.DB.Setting.Timeout)*time.Second)
	defer cancel()

	db, err := sqlx.ConnectContext(ctxTimeout, "postgres", getConnection())
	if err != nil {
		log.Fatalln("[DB - Init] Error initializing database, err: ", err)
		return err
	}

	if err := db.Ping(); err != nil {
		log.Fatalln("[DB - Init] Error pinging database, err: ", err)
	}

	db.SetConnMaxIdleTime(time.Duration(config.Configs.DB.Setting.MaxIdleTime) * time.Second)
	db.SetConnMaxLifetime(time.Duration(config.Configs.DB.Setting.MaxLifetime) * time.Second)
	db.SetMaxIdleConns(config.Configs.DB.Setting.MaxIdleConns)
	db.SetMaxOpenConns(config.Configs.DB.Setting.MaxOpenConns)

	DB = db

	return nil
}
