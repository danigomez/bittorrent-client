package torrent

/* BEP 15 */

// 16 bytes
type ConnectRequest struct {
	protocolId    int64 // offset: 0
	action        int32 // offset: 8
	transactionId int32 // offset: 12
}

// 16 bytes
type ConnectResponse struct {
	action        int32 // offset: 0
	connectionId  int64 // offset: 8
	transactionId int32 // offset: 4
}

// 98 bytes
type AnnounceRequest struct {
	connectionId  int64    // offset: 0
	action        int32    // offset: 8 1: announce
	transactionId int32    // offset: 12
	infoHash      [20]byte // offset: 16
	peerId        [20]byte // offset: 36
	downloaded    int64    // offset: 56
	left          int64    // offset: 64
	uploaded      int64    // offset: 72
	event         int32    // offset: 80  0: none; 1: completed; 2: started; 3: stopped
	ipAddress     int32    // offset: 84 0 default
	key           int32    // offset: 88
	numWant       int32    // offset: 92 -1 default
	port          int16    // offset: 96
}

type AnnounceResponse struct {
	action        int32 // 1: announce
	transactionId int32
	interval      int32
	leechers      int32
	seeders       int32
	addresses     []peerAddress
}

type peerAddress struct {
	ipAddress int32
	tcpPort   int16
}
