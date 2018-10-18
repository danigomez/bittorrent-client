package peer

import (
	"github.com/danigomez/bittorrent-client/torrent"
	"github.com/danigomez/bittorrent-client/torrent/file"
	"github.com/danigomez/bittorrent-client/torrent/tracker"
	"log"
	"testing"
)

func TestConnerToPeer(t *testing.T) {

	f := file.Open("../test.torrent")
	trackerUrl := "9.rarbg.to:2710"
	connect, err := tracker.SendConnectRequest(trackerUrl)

	if err != nil {
		t.Errorf("error: there was an error while sending connect request %s", err)

	}

	connectionId := connect.ConnectionId
	hash, _ := f.GetInfoHash()
	id := torrent.GetId("-BBV001-")
	size := f.GetSize()

	_, err = tracker.SendAnnounceRequest(trackerUrl, connectionId, hash, id, size)

	if err != nil {
		log.Fatalf("%s", err)
		t.Errorf("%s", err)
	}

}
