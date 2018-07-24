package torrent

import (
	"log"
	"math/rand"
	"time"
)

const maxIdSize = 20

func GetId(clientName string) []byte {

	length := len(clientName)
	if length > maxIdSize {
		log.Panicf(" Wrong length for client name %v ", length)
	}

	rand.Seed(time.Now().UnixNano())

	id := make([]byte, maxIdSize)

	rand.Read(id)

	return append([]byte(clientName), id[length:]...)
}
