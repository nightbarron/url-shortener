package services

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"math"
	"url-shortener/models"
)

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
