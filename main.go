package main

import (
	"Exoplanet/ApiHandler"
	"Exoplanet/Database"
	"Exoplanet/Domain"
	"Exoplanet/Models"
	"fmt"
	"os"
)

func loadEnv() {
	Models.DbUser = os.Getenv("DB_USER")
	Models.DbPassword = os.Getenv("DB_PASSWORD")
	Models.DbHost = os.Getenv("DB_HOST")
	Models.DbPort = os.Getenv("DB_PORT")
	Models.DbName = os.Getenv("DB_NAME")
	Models.AppPort = os.Getenv("APP_PORT")
}

func main() {
	loadEnv()
	newDBService := Database.NewDBService()
	domainService := Domain.NewDomainService(newDBService)
	webService := ApiHandler.NewWebservices(fmt.Sprintf(":%s", Models.AppPort), domainService)
	err := webService.Start()
	if err != nil {
		return
	}
}
