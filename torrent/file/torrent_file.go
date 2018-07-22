package file

import (
	"fmt"
	"github.com/zeebo/bencode"
	"io/ioutil"
)

type File struct {
	Announce     string `bencode:"announce"`
	CreatedBy    string `bencode:"created by"`
	CreationDate int    `bencode:"creation date"`
	Encoding     string `bencode:"encoding"`
	Info         info   `bencode:"info"`
}

type info struct {
	Length      int    `bencode:"length"`
	Name        string `bencode:"name"`
	PieceLength int    `bencode:"piece length"`
	Pieces      string `bencode:"pieces"`
}

func Open(fileName string) File {
	file, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Errorf("error: there was an error while opening torrent %s", err)
	}

	return parseTorrentData(file)
}

func OpenFromString(torrent string) File {
	return parseTorrentData([]byte(torrent))
}

func parseTorrentData(torrent []byte) File {
	var parsed File

	// Parse bencode torrent file
	err := bencode.DecodeBytes(torrent, &parsed)

	if err != nil {
		fmt.Errorf("error: there was an error parsing torrent data %s", err)
	}

	return parsed

}
