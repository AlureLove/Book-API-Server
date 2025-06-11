package main

import (
	"Book-API-Server/api"
	"Book-API-Server/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = "config/application.yaml"
	}
	if err := config.LoadConfigFromYaml(path); err != nil {
		fmt.Printf("load config err: %s\n", err)
		os.Exit(1)
	}

	conf := config.Get()

	server := gin.Default()

	api.NewBookApiHandler().Registry(server)

	if err := server.Run(conf.App.Address()); err != nil {
		log.Println(err)
	}
}
