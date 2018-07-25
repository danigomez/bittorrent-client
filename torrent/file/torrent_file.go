package file

import (
	"crypto/sha1"
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
	Files       []files `bencode:"files,omitempty"`
	Length      int     `bencode:"length,omitempty"`
	Name        string  `bencode:"name,omitempty"`
	PieceLength int     `bencode:"piece length,omitempty"`
	Pieces      string  `bencode:"pieces,omitempty"`
}

type files struct {
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

func (file File) GetInfoHash() ([20]byte, error) {
	rawInfo, err := bencode.EncodeBytes(file.Info)

	if err != nil {
		return [20]byte{0}, fmt.Errorf("error: there was an error while getting info data")
	}

	return sha1.Sum(rawInfo), nil
}

func (file File) GetSize() int64 {
	infoData := file.Info

	if infoData.Length > 0 {
		return int64(infoData.Length)
	} else {
		var size int64 = 0
		for _, file := range infoData.Files {
			size += int64(file.Length)
		}
		return size
	}

}
