package main

import (
	"url-shortener/configs"
	"url-shortener/helpers"
	route "url-shortener/routes"
)

// "hash"

func main() {

	// Init log
	helpers.InitLogger()

	// Init config
	config := configs.GlobalConfig{}
	err := config.Load("configs/config.json")
	if err != nil {
		return
	}

	// Salt for hashing
	// @deprecated
	var saltList []string

	r := route.SetupRouter(config, &saltList)
	err = r.Run(":8080")
	if err != nil {
		return
	}

}
