package config

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"

	"gogenggo/internals/types/constants"

	"gopkg.in/ini.v1"
)

var Configs *Config

func LoadConfig() error {
	if Configs != nil {
		return nil
	}

	env := os.Getenv(constants.GolangChatbotEnvironment)
	if env == "" {
		env = constants.DevelopmentEnvironment
	}

	mainConfig := MainConfig{}
	var mainConfigFile any = fmt.Sprintf(constants.MainConfigFile, env)
	if _, err := os.Stat(fmt.Sprintf(constants.MainConfigFile, env)); err != nil {
		if os.Getenv(constants.APIConfigPath) == "" {
			log.Fatalln("[Config - LoadConfig] Error loading config file from url, environment variable empty")
		}

		resp, _ := http.Get(fmt.Sprintf("%s/main.%s-config.ini", os.Getenv(constants.APIConfigPath), env))
		defer resp.Body.Close()

		mainConfigFile = resp.Body
	}

	if err := ini.MapTo(&mainConfig, mainConfigFile); err != nil {
		log.Fatal("[Config - LoadConfig] Error loading and mapping main config, err:", err)
		return err
	}

	if !reflect.DeepEqual(Configs.Main, mainConfig) {
		log.Printf("Main config changed, values: %#v", mainConfig)
	}

	dbConfig := DBConfig{}
	var dbConfigFile any = fmt.Sprintf(constants.DBConfigFile, env)
	if _, err := os.Stat(fmt.Sprintf(constants.DBConfigFile, env)); err != nil {
		if os.Getenv(constants.APIConfigPath) == "" {
			log.Fatalln("[Config - LoadConfig] Error loading config file from url, environment variable empty")
		}

		resp, _ := http.Get(fmt.Sprintf("%s/db.%s-config.ini", os.Getenv(constants.APIConfigPath), env))
		defer resp.Body.Close()

		dbConfigFile = resp.Body
	}

	if err := ini.MapTo(&dbConfig, dbConfigFile); err != nil {
		log.Fatal("[Config - LoadConfig] Error loading and mapping db config, err:", err)
		return err
	}

	if !reflect.DeepEqual(Configs.DB, dbConfig) {
		log.Printf("DB config changed, values: %#v", dbConfig)
	}

	messageConfig := MessageConfig{}
	var messageConfigFile any = fmt.Sprintf(constants.MessageConfigFile, env)
	if _, err := os.Stat(fmt.Sprintf(constants.MessageConfigFile, env)); err != nil {
		if os.Getenv(constants.APIConfigPath) == "" {
			log.Fatalln("[Config - LoadConfig] Error loading config file from url, environment variable empty")
		}

		resp, _ := http.Get(fmt.Sprintf("%s/message.%s-config.ini", os.Getenv(constants.APIConfigPath), env))
		defer resp.Body.Close()

		messageConfigFile = resp.Body
	}

	if err := ini.MapTo(&messageConfig, messageConfigFile); err != nil {
		log.Fatal("[Config - LoadConfig] Error loading and mapping message config, err:", err)
		return err
	}

	if !reflect.DeepEqual(Configs.Message, messageConfig) {
		log.Printf("Message config changed, values: %#v", messageConfig)
	}

	Configs = &Config{
		Main:    mainConfig,
		DB:      dbConfig,
		Message: messageConfig,
	}

	return nil
}
