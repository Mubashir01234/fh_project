package main

import (
	"log"
	"net/http"
	"os"
	"project/handlers"
	"project/routes"

	"github.com/fatih/color"
)

func main() {
	logFile, err := os.OpenFile(handlers.DotEnvVariable("LOG_FILE_PATH"), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("unable to open csv file: %v\n", err)
		return
	}
	log.SetOutput(logFile)
	color.Cyan("🌏 Server running on localhost:" + handlers.DotEnvVariable("SERVER_PORT"))
	log.Fatal(http.ListenAndServe(":"+handlers.DotEnvVariable("SERVER_PORT"), routes.Routes()))
}
