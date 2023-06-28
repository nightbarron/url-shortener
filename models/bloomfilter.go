package models

import (
	"hash"
	"math"
	"url-shortener/helpers"
)

type BloomFilters interface {
	Set(value string)
	Check(value string) bool
}
type bloomFilterType struct {
	bitfield     []bool
	hashFuncList []hash.Hash
	size         uint64
}

// CreateBloomFilter Create bloom filter
func CreateBloomFilter(hashList []hash.Hash, bitSetSize uint64) BloomFilters {
	return &bloomFilterType{
		bitfield:     make([]bool, bitSetSize),
		hashFuncList: hashList,
		size:         bitSetSize,
	}
}

// Set value for each element
func (b *bloomFilterType) Set(value string) {
	for _, h := range b.hashFuncList {
		hashPosition := getPosition(helpers.Hashing(h, value), b.size)
		b.bitfield[hashPosition] = true
	}
}

func (b *bloomFilterType) Check(value string) bool {
	for _, h := range b.hashFuncList {
		return b.bitfield[getPosition(helpers.Hashing(h, value), b.size)]
	}
	return true
}

// Return the value associate with each element in filter array
func getPosition(hashValue int64, bitSetSize uint64) int64 {
	if hashValue < 0 {
		hashValue = -hashValue
	}
	return int64(math.Abs(float64(uint64(hashValue) % bitSetSize)))
}
