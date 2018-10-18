package peer

import (
	"fmt"
	"github.com/danigomez/bittorrent-client/torrent/network/broker"
	"github.com/danigomez/bittorrent-client/torrent/network/peer"
)

func SendHandshakeRequest(infoHash [20]byte, peerId [20]byte, peerAddress string) ([]byte, error) {
	brokerClient := new(broker.TCPBroker)
	data, err := peer.NewHandshakeRequest(infoHash, peerId).Serialize()

	if err != nil {
		return nil, fmt.Errorf("error: there was an error while serializing data %s, \n%s", peerAddress, err)
	}

	// Creates new request
	request := broker.NewBrokerRequest(peerAddress, data)
	response, err := brokerClient.SendRequest(request)

	if err != nil {
		return nil, fmt.Errorf("error: there was an error while serializing data %s, \n%s", peerAddress, err)
	}

	return response, nil
}
