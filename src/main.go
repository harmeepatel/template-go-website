package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"makeyourcareer.net/app"
)

func init() {
    const env = "../.env"
	switch depenv := os.Getenv("DEPLOY_PLATFORM"); depenv {
	case "prod":
		log.Println("-- Production evironment --")
		if godotenv.Load(env) != nil {
			log.Fatal(".env not found")
		}
	case "test":
		log.Println("-- Testing evironment --")
		if godotenv.Load(env + ".test") != nil {
			log.Fatal(".env.test not found")
		}
	default:
		log.Println("-- dev --")
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
