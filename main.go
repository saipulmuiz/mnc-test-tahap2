package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
	"github.com/saipulmuiz/mnc-test-tahap2/config"
	"github.com/saipulmuiz/mnc-test-tahap2/routers"
)

func main() {
	log.SetReportCaller(true)

	//Load Env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.ConnectDB()

	route := routers.RouterConfig(db)

	route.Run(os.Getenv("APP_PORT"))
}
