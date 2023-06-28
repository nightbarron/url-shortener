package main

import (
	"url-shortener/configs"
	"url-shortener/helpers"
	"url-shortener/models"
	route "url-shortener/routes"
)

// "hash"

func main() {

	// Init log
	helpers.InitLogger()

	// Init bloom filter
	bloomFilter := models.InitBloomFilter()

	// REFILl short url to bloom filter
	// TODO: implement this function

	// Init config
	config := configs.GlobalConfig{}
	err := config.Load("configs/config.json")
	if err != nil {
		return
	}

	// Salt for hashing
	// @deprecated
	//var saltList []string

	r := route.SetupRouter(config, &bloomFilter)
	err = r.Run(":8080")
	if err != nil {
		return
	}

}
