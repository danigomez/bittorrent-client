package tracker

import (
	"fmt"
	"github.com/danigomez/bittorrent-client/torrent/network/broker"
	"github.com/danigomez/bittorrent-client/torrent/network/tracker"
	"net/url"
	"strconv"
)

func SendConnectRequest(trackerUrl string) (*tracker.ConnectResponse, error) {

	brokerClient := new(broker.UDPBroker)
	data, err := tracker.NewConnectRequest().Serialize()

	if err != nil {
		return nil, fmt.Errorf("error: there was an error while serializing data %s, \n%s", trackerUrl, err)
	}

	// Creates new request
	request := broker.NewBrokerRequest(trackerUrl, data)
	response, err := brokerClient.SendRequest(request)

	if err != nil {
		return nil, fmt.Errorf("error: there was an error sending connect request to tracker %s, \n%s", trackerUrl, err)
	}

	ret := new(tracker.ConnectResponse)

	err = ret.Deserialize(response)

	if err != nil {
		return nil, fmt.Errorf("error: there was an error while deserializing data %s, \n%s", trackerUrl, err)
	}

	return ret, nil

}

func SendAnnounceRequest(trackerUrl string, connectionId int64, infoHash [20]byte, peerId [20]byte, size int64) (*tracker.AnnounceResponse, error) {

	brokerClient := new(broker.UDPBroker)

	parsedUrl, err := url.Parse(trackerUrl)

	if err != nil {
		return nil, fmt.Errorf("error: there was an error while parsing tracker URL %s, \n%s", trackerUrl, err)
	}

	port, _ := strconv.Atoi(parsedUrl.Port())

	data, err := tracker.NewAnnounceRequest(connectionId, infoHash, peerId, size, int16(port)).Serialize()

	if err != nil {
		return nil, fmt.Errorf("error: there was an error while serializing data %s, \n%s", trackerUrl, err)
	}

	// Creates new request
	request := broker.NewBrokerRequest(trackerUrl, data)
	response, err := brokerClient.SendRequest(request)

	if err != nil {
		return nil, fmt.Errorf("error: there was an error sending announce request to tracker %s, \n%s", trackerUrl, err)
	}
	ret := new(tracker.AnnounceResponse)

	err = ret.Deserialize(response)

	if err != nil {
		return nil, fmt.Errorf("error: there was an error while deserializing data %s, \n%s", trackerUrl, err)
	}

	return ret, nil

}
