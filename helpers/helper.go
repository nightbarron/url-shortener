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

func ReadFile(path string) (string, []byte) {
	log.Info("Reading file: ", path)
	// Read content from file
	var file []byte = nil
	var err error
	file, err = os.ReadFile(path)
	//os.Stdout.Write(file)
	if err != nil {
		log.Error("Fail to read in ", path)
		return "Fail to read in " + path, nil
	}
	return string(file), file
}
