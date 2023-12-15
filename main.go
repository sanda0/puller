package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sanda0/puller/handlers"
	"github.com/sanda0/puller/setup"
)

func server() {
	server := gin.Default()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081" // Default port if not specified
	}

	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	server.POST("/puller", handlers.HandleWebHook)

	server.Run(":" + port)
}

func main() {
	startListener := flag.Bool("s", false, "Start webhook listener")
	init := flag.Bool("i", false, "Initialize config")
	flag.Parse()

	if *startListener {
		server()
	}

	if *init {
		fmt.Println("Initializing service...")
		err := setup.CreateServiceFile()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			fmt.Println("Service created successfully")
		}
		fmt.Println("Creating config.json")
		err = setup.GenerateConfigFile()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			fmt.Println("Config created successfully")
		}
	}
}
