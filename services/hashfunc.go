package services

import (
	"hash/fnv"
	"math/rand"
	"strconv"
	"time"
	"url-shortener/configs"
)

func StringHashToNumber(input string, k int) string {
	// This function will return a number from 1 to k
	hashed := fnv.New32()
	_, err := hashed.Write([]byte(input))
	if err != nil {
		// Default to 1
		return "1"
	}
	hashValue := hashed.Sum32()
	outputNumber := int(hashValue%uint32(k)) + 1
	return strconv.Itoa(outputNumber)
}

func AppendSaltToString(originStr string) string {
	// Get random salt
	salt := generateSalt(5)
	return originStr + salt
}

func generateSalt(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	salt := make([]byte, length)
	for i := range salt {
		salt[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(salt)
}

func ISDuplicateShortUrl(globalConfig configs.GlobalConfig, shortUrl string) bool {
	// Check if shortUrl is already in database
	// su dung bloom filter

	return false
}
