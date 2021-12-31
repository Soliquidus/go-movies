package main

import (
	"log"
	"server/cmd/api"
)

func main() {
	api.Open()
	defer api.Close()
	router := api.SetupRoutes()
	err := router.Run(":4000")
	if err != nil {
		log.Println(err)
	}
}

