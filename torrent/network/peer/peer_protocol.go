package peer

import (
	"bytes"
	"encoding/binary"
	"github.com/danigomez/bittorrent-client/torrent"
)

type HandshakeRequest struct {
	pstrlen  int32
	pstr     string
	reserved int64
	infoHash [20]byte
	peerId   [20]byte
}

func NewHandshakeRequest(infoHash [20]byte, peerId [20]byte) HandshakeRequest {
	return HandshakeRequest{
		int32(torrent.BitTorrentPstrLen),
		torrent.BitTorrentPstr,
		0,
		infoHash,
		peerId,
	}
}

func (hsr HandshakeRequest) Serialize() ([]byte, error){
	buffer := new(bytes.Buffer)

	err := binary.Write(buffer, binary.BigEndian, hsr)

	return buffer.Bytes(), err
}