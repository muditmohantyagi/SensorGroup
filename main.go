package main

import (
	"os"

	_ "test.com/docs"
	"test.com/pkg/cmd"
	"test.com/router"
)

// @title Sensor data service API
// @version 1.0
// @description This api witll generate sencer data related to fish.
// @host localhost:8080
// @BasePath /api

func main() {
	r := router.SetupRouter()
	args := os.Args

	if len(args) > 1 {
		cmd.Execute()
		os.Exit(1)
	}

	r.Run(":8080")
}
