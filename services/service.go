package services

import (
	"encoding/binary"
	"hash/crc32"
	"strconv"
	"url-shortener/configs"
	"url-shortener/models"
)

var bloomFilter models.BloomFilters

func IsExistByBoomFilter(config configs.GlobalConfig, longUrl string) bool {
	return false
}

func GenShortUrlHash(config configs.GlobalConfig, longUrl string) (string, error) {
	// CRC32 here
	data := []byte(longUrl)
	hashFunc := crc32.NewIEEE()

	_, err := hashFunc.Write(data)
	if err != nil {
		return "", err
	}
	checksum := hashFunc.Sum32()

	// Convert uint32 to byte slice
	byteSlice := make([]byte, 4)
	binary.BigEndian.PutUint32(byteSlice, checksum)

	asciiChecksum := strconv.FormatUint(uint64(checksum), 36)

	return asciiChecksum, nil
}

func LookUpLongUrl(config configs.GlobalConfig, shortUrl string) (string, error) {
	return "", nil
}

func IsMappedLongToShortDB(config configs.GlobalConfig, longUrl string) bool {
	//TODO: implement this function
	// Check in DB
	return false
}
