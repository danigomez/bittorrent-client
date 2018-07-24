package torrent

import (
	"fmt"
	"github.com/danigomez/bittorrent-client/torrent/network/broker"
	"github.com/danigomez/bittorrent-client/torrent/network/tracker"
)

func ConnectToTracker(trackerUrl string) (*tracker.ConnectResponse, error) {

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

/*

func AnnounceToTracker(trackerUrl string) (*tracker.AnnounceResponse, error) {

	brokerClient := new(broker.UDPBroker)
	data, err := tracker.NewAnnounceRequest().Serialize()

	if err != nil {
		return nil, fmt.Errorf("error: there was an error while serializing data %s, \n%s", trackerUrl, err)
	}

	// Creates new request
	request := broker.NewBrokerRequest(trackerUrl, data)
	_, err = brokerClient.SendRequest(request)

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
*/
