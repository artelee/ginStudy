package main

import (
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
	config.DBClient()
	router.GetCommonRouter(8080)
}
