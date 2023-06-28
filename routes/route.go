package routes

import (
	"github.com/gin-gonic/gin"
	"url-shortener/configs"
	ctl "url-shortener/controllers"
	"url-shortener/models"
)

type Routes struct {
	router *gin.Engine
}

func SetupRouter(config configs.GlobalConfig, bloomFilter *models.BloomFilters) *gin.Engine {

	r := Routes{
		router: gin.Default(),
	}

	r.router.Use(func(c *gin.Context) {
		// add header Access-Control-Allow-Origin
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	})

	//API route for version 11
	apiV1 := r.router.Group("/v1/api/url-shorten")
	apiV1.GET("/version", ctl.GetVersion())

	apiV1.POST("/genshorturl", ctl.GenShortUrl(config, bloomFilter))

	apiV1.GET("/getlongurl", ctl.GetLongUrl(config))

	return r.router

}
