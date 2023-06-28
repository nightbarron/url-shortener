package helpers

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash"
	"os"

	log "github.com/sirupsen/logrus"
)

func InitLogger() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func ToString(value interface{}) string {
	return fmt.Sprintf("%v", value)
}

func Hashing(h hash.Hash, value string) int64 {
	h.Write([]byte(value))
	bits := h.Sum(nil)
	buffer := bytes.NewBuffer(bits)
	result, _ := binary.ReadVarint(buffer)
	h.Reset()
	return result
}
