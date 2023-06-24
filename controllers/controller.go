package controllers

import (
	// svc "gin_template/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

var storageFolderPath = "/etc/api/data/"

func GetVersion() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"version": "1.0.0"})
	}
}

// func DeleteKey() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		key := c.Param("key")
// 		err := svc.DeleteFile(storageFolderPath + key)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		} else {
// 			c.JSON(http.StatusOK, gin.H{"message": "success"})
// 		}
// 	}
// }
