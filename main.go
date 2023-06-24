package main

import (
	"gin_template/services"
)

// "hash"

func main() {
	services.BoomFilterApply()
}

// func createIntHash(input string) int {
// 	bits := hasher.Sum([]byte(input))
// 	buf := bytes.NewBuffer(bits)
// 	result, _ := binary.ReadVarint(buf)
// 	return int(result)
// }

// func createNormalHash(input string) {
// 	fmt.Printf("%x\n", hasher.Sum([]byte(input)))
// }

// func (f *filter) hashPosition(s string) int {
// 	hs := createIntHash(s)
// 	if hs < 0 {
// 		hs = -hs
// 	}
// 	return hs % len(f.bitfield)
// }

// func (f *filter) set(s string) {
// 	tmp := f.hashPosition(s)
// 	f.bitfield[tmp] = true
// }
// func (f *filter) get(s string) bool {
// 	return f.bitfield[f.hashPosition(s)]
// }
