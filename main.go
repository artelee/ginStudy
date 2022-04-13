package main

import (
	"ginStudy/api"
	"ginStudy/config"
	"ginStudy/router"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	client := config.DBClient()
	api.GetUserList(client)
	router.GetCommonRouter(8080)
}
