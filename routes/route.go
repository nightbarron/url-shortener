package routes

// import (
// 	ctl "gin_template/controllers"
// 	"github.com/gin-gonic/gin"
// )

// type Routes struct {
// 	router *gin.Engine
// }

// func SetupRouter() *gin.Engine {

// 	r := Routes{
// 		router: gin.Default(),
// 	}

// 	r.router.Use(func(c *gin.Context) {
// 		// add header Access-Control-Allow-Origin
// 		c.Writer.Header().Set("Content-Type", "application/json")
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, UPDATE")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
// 		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(200)
// 		} else {
// 			c.Next()
// 		}
// 	})

// 	//API route for version 1
// 	apiV1 := r.router.Group("/v1/api/template")
// 	apiV1.GET("/version", ctl.GetVersion())
// 	apiV1.DELETE("/key/:key", ctl.DeleteKey())

// 	return r.router

// }
