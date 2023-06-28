package controllers

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"

	// svc "gin_template/services"

	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
	"url-shortener/configs"
	"url-shortener/models"
	svc "url-shortener/services"
)

func GetVersion() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"version": "1.0.0"})
	}
}

func GenShortUrl(globalConfig configs.GlobalConfig, filters *models.BloomFilters) gin.HandlerFunc {
	return func(c *gin.Context) {
		var urlRequest models.UrlRequest
		var shortUrlHash string
		// Get request body
		body, err := ioutil.ReadAll(c.Request.Body)
		ip := c.Request.RemoteAddr

		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Error("Error when close body: ", err)
				return
			}
		}(c.Request.Body)

		if err != nil {
			log.Error("Error when read body: ", err)
			_ = c.AbortWithError(400, err)
			return
		}
		err = json.Unmarshal(body, &urlRequest)
		if err != nil {
			log.Error("Error when unmarshal body: ", err)
			_ = c.AbortWithError(400, err)
			return
		}

		longUrl := urlRequest.Url + ip

		// Check if longUrl is already in database
		// => Khong nen check vi se lam overload database

		// Generate short url
		shortUrlHash, err = svc.GenShortUrlHash(globalConfig, longUrl)

		for svc.ISDuplicateShortUrl(globalConfig, shortUrlHash, filters) {
			longUrl = svc.AppendSaltToString(longUrl)
			shortUrlHash, err = svc.GenShortUrlHash(globalConfig, longUrl)
		}

		// Add short url to bloom filters
		(*filters).Set(shortUrlHash)

		// WORKING WITH DATABASES
		// Get database name
		dbPostFix := svc.StringHashToNumber(shortUrlHash, globalConfig.MySQL.NumberOfDB)
		err = svc.SaveLongShortToDB(globalConfig, longUrl, shortUrlHash, globalConfig.MySQL.DBNamePrefix+dbPostFix)
		if err != nil {
			log.Error("Error when save long short url to database: ", err)
			c.JSON(http.StatusOK, models.UrlResponse{Success: false, Url: ""})
			return
		}

		c.JSON(http.StatusOK, models.UrlResponse{Success: true, Url: globalConfig.ShortUrlDomain + shortUrlHash})

	}
}

func GetLongUrl(globalConfig configs.GlobalConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: implement
		// The best way is to curl to nginx to get long url

		var urlRequest models.UrlRequest
		// Get request body
		body, err := ioutil.ReadAll(c.Request.Body)

		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Error("Error when close body: ", err)
				return
			}
		}(c.Request.Body)

		if err != nil {
			log.Error("Error when read body: ", err)
			_ = c.AbortWithError(400, err)
			return
		}
		err = json.Unmarshal(body, &urlRequest)
		if err != nil {
			log.Error("Error when unmarshal body: ", err)
			_ = c.AbortWithError(400, err)
			return
		}

		shortUrl := urlRequest.Url

		longUrl, err := svc.LookUpLongUrl(globalConfig, shortUrl)
		if err != nil {
			log.Error("Error when lookup long url: ", err)
			_ = c.AbortWithError(500, err)
			return
		}

		c.JSON(http.StatusOK, models.UrlResponse{Success: true, Url: longUrl})

	}
}
