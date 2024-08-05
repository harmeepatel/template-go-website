package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/harmeepatel/template-go-website/app"
)

func init() {
    const env = "../.env"
	switch depenv := os.Getenv("DEPLOY_PLATFORM"); depenv {
	case "prod":
		log.Println("-- PRODUCTION --")
		if godotenv.Load(env) != nil {
			log.Fatal(".env not found")
		}
	default:
		log.Println("-- DEV --")
		if godotenv.Load(env + ".dev") != nil {
			log.Fatal(".env.dev not found")
		}
	}
}

func main() {
	log.Println("Listening on: " + os.Getenv("PORT"))
	log.Println("pid         :", os.Getpid())

	mainServer := app.NewApp(os.Getenv("PORT"))
	mainServer.InitStaticServer()
	mainServer.InitRoutes()
	mainServer.Run()
}
