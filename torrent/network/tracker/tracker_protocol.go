package tracker

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand"
	"time"
)

/* BEP 15 */

const protocolId = 0x41727101980

const (
	actionConnect = iota
	actionAnnounce
)

const (
	eventNone = iota
	eventCompleted
	eventStarted
)

// 16 bytes
type ConnectRequest struct {
	ProtocolId    int64 // offset: 0
	Action        int32 // offset: 8
	TransactionId int32 // offset: 12
}

// 16 bytes
type ConnectResponse struct {
	Action        int32 // offset: 0
	TransactionId int32 // offset: 4
	ConnectionId  int64 // offset: 8
}

// 98 bytes
type AnnounceRequest struct {
	ConnectionId  int64    // offset: 0
	Action        int32    // offset: 8 1: announce
	TransactionId int32    // offset: 12
	InfoHash      [20]byte // offset: 16
	PeerId        [20]byte // offset: 36
	Downloaded    int64    // offset: 56
	Left          int64    // offset: 64
	Uploaded      int64    // offset: 72
	Event         int32    // offset: 80  0: none; 1: completed; 2: started; 3: stopped
	IpAddress     int32    // offset: 84 0 default
	Key           int32    // offset: 88
	NumWant       int32    // offset: 92 -1 default
	Port          int16    // offset: 96
}

type AnnounceResponse struct {
	Action        int32 // 1: announce
	TransactionId int32
	Interval      int32
	Leechers      int32
	Seeders       int32
	Addresses     []peerAddress
}

type peerAddress struct {
	IpAddress int32
	TcpPort   int16
}

func NewConnectRequest() ConnectRequest {
	rand.Seed(time.Now().UnixNano())
	return ConnectRequest{
		protocolId,
		actionConnect,
		rand.Int31(),
	}
}

func NewAnnounceRequest(connectionId int64, infoHash [20]byte, peerId [20]byte, size int64, port int16) AnnounceRequest {
	rand.Seed(time.Now().UnixNano())
	return AnnounceRequest{
		connectionId,
		actionAnnounce,
		rand.Int31(),
		infoHash, // SHA-1 of info dict from torrent file
		peerId,
		0,
		size,
		0,
		eventNone,
		0,
		rand.Int31(),
		-1,
		port,
	}
}

func (cr ConnectRequest) Serialize() ([]byte, error) {
	buffer := new(bytes.Buffer)

	err := binary.Write(buffer, binary.BigEndian, cr)

	return buffer.Bytes(), err
}

func (cr *ConnectResponse) Deserialize(data []byte) error {
	buf := bytes.NewReader(data)
	err := binary.Read(buf, binary.BigEndian, cr)

	return err
}

func (ar AnnounceRequest) Serialize() ([]byte, error) {
	buffer := new(bytes.Buffer)

	err := binary.Write(buffer, binary.BigEndian, ar)

	return buffer.Bytes(), err
}

func (ar *AnnounceResponse) Deserialize(data []byte) error {
	dataLen := len(data)
	buf := bytes.NewReader(data)

	err := binary.Read(buf, binary.BigEndian, &ar.Action)

	if err != nil {
		return fmt.Errorf("error: There was an error deserializing Action field \n%s", err)
	}

	err = binary.Read(buf, binary.BigEndian, &ar.TransactionId)

	if err != nil {
		return fmt.Errorf("error: There was an error deserializing TransactionId field \n%s", err)
	}

	err = binary.Read(buf, binary.BigEndian, &ar.Interval)

	if err != nil {
		return fmt.Errorf("error: There was an error deserializing Interval field \n%s", err)
	}

	err = binary.Read(buf, binary.BigEndian, &ar.Leechers)

	if err != nil {
		return fmt.Errorf("error: There was an error deserializing Leechers field \n%s", err)
	}

	err = binary.Read(buf, binary.BigEndian, &ar.Seeders)

	if err != nil {
		return fmt.Errorf("error: There was an error deserializing Seeders field \n%s", err)
	}

	if dataLen > 20 {
		extraBuf := bytes.NewReader(data[20:])

		// If there is more that 20 bytes, then there is peer data,
		// Each peer is grouped in a 6 byte block
		peersCount := (dataLen - 20) / 6

		for i := 0; i < peersCount; i++ {
			var address int32
			var port int16
			err = binary.Read(extraBuf, binary.BigEndian, &address)

			if err != nil {
				return fmt.Errorf("error: There was an error deserializing Address field \n%s", err)
			}

			err = binary.Read(extraBuf, binary.BigEndian, &port)
			if err != nil {
				return fmt.Errorf("error: There was an error deserializing Port field \n%s", err)

			}
			ar.Addresses = append(ar.Addresses, peerAddress{address, port})

		}
	}

	return nil
}
