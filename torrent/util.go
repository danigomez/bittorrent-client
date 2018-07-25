package torrent

import (
	"log"
	"math/rand"
	"time"
)

const maxIdSize = 20

func GetId(clientName string) [20]byte {

	var ret [20]byte

	length := len(clientName)
	if length > maxIdSize {
		log.Panicf(" Wrong length for client name %v ", length)
	}

	rand.Seed(time.Now().UnixNano())

	id := make([]byte, maxIdSize)

	rand.Read(id)

	copy(ret[:], append([]byte(clientName), id[length:]...))

	return ret
}
