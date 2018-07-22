package torrent

import (
	"github.com/danigomez/bittorrent-client/torrent/network/tracker"
	"log"
	"net"
)

func ConnectToTracker(trackerUrl string) *tracker.ConnectResponse {

	connection, err := net.Dial("udp", trackerUrl)

	defer connection.Close()

	if err != nil {
		log.Fatalf("error: there was an error while creating connection to %s, \n%s", trackerUrl, err)
	}

	request := tracker.NewConnectRequest()

	encoded, err := request.Serialize()

	_, err = connection.Write(encoded)

	if err != nil {
		log.Fatalf("error: there was an error sending UDP data to tracker \n%s", err)
	}

	buffer := make([]byte, 2048)

	_, err = connection.Read(buffer)

	if err != nil {
		log.Fatalf("error: there was an error while reading UDP data from %s, \n%s", trackerUrl, err)
	}

	ret := new(tracker.ConnectResponse)

	err = ret.Deserialize(buffer)

	if err != nil {
		log.Fatalf("error: there was an error while deserializing data %s, \n%s", trackerUrl, err)
	}

	return ret

}
