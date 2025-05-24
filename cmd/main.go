package main

import (
	"croner-scheduler/internal/config"
	"croner-scheduler/internal/database"
	"croner-scheduler/internal/initializers"
	"croner-scheduler/internal/scheduler"
	"log"
)

func main() {
	initializers.LoadEnvVariables()
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	database.NewDB(cfg)
	defer database.CloseDB()

	if err := scheduler.Run(cfg); err != nil {
		log.Fatal(err)
	}
}
