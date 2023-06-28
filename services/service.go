package services

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/binary"
	"fmt"
	"hash"
	"hash/crc32"
	"math"
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
	hash := crc32.NewIEEE()

	_, err := hash.Write(data)
	if err != nil {
		return "", err
	}
	checksum := hash.Sum32()

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

func BoomFilterApply() {
	bloomFilter = initBloomFilter()
	var s1 = "hoale"
	bloomFilter.Set(s1)
	fmt.Println(bloomFilter.Check(s1))
	fmt.Println(bloomFilter.Check("hoalethe"))

}

func initBloomFilter() models.BloomFilters {
	hashFunctions := []hash.Hash{
		sha256.New224(),
		sha256.New(),
		sha512.New(),
		sha512.New384(),
	}

	bloomFilter = models.CreateBloomFilter(hashFunctions, math.MaxUint32)
	return bloomFilter
}
