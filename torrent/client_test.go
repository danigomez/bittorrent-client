package torrent

import (
	"github.com/danigomez/bittorrent-client/torrent/file"
	"reflect"
	"testing"
)

func TestConnectToTracker(t *testing.T) {
	data, _ := ConnectToTracker("tracker.internetwarriors.net:1337")

	if reflect.TypeOf(data.TransactionId).String() != "int32" {
		t.Errorf("error: TransactionId is not int32")
	}

	if reflect.TypeOf(data.Action).String() != "int32" {
		t.Errorf("error: Action is not int32")
	}

	if reflect.TypeOf(data.ConnectionId).String() != "int64" {
		t.Errorf("error: ConnectionId is not int64")
	}

	if data.Action != 0 {
		t.Errorf("error: Action is not equal to 0 (CONNECT)")
	}

	_, err := ConnectToTracker("satasa.com")

	if err == nil {
		t.Errorf("error: There should be an error when connecting to invalid URI")
	}

}

func TestAnnounceToTracker(t *testing.T) {

	f := file.OpenFromString(`d8:announce43:udp://tracker.coppersurfer.tk:6969/announce10:created by13:uTorrent/187013:creation datei1462355939e8:encoding5:UTF-84:infod6:lengthi124234e4:name9:puppy.jpg12:piece lengthi16384e6:pieces160:T�k�/�_(�S\u0011h%���+]q\'B\u0018�٠:����p"�j����1-g"\u0018�s(\u001b\u000f���V��=�h�m\u0017a�nF�2���N\r�ǩ�_�\u001e"2���\'�wO���-;\u0004ע\u0017�ؑ��L&����0\u001f�D_9��\t\\��O�h,n\u001a5g�(��仑,�\\߰�%��U��\u0019��C\u0007>��df��ee`)
	trackerUrl := "tracker.coppersurfer.tk:6969"
	connect, _ := ConnectToTracker(trackerUrl)

	connectionId := connect.ConnectionId
	hash, _ := f.GetInfoHash()
	id := GetId("-BBV001-")
	size := f.GetSize()

	ret, err := AnnounceToTracker(trackerUrl, connectionId, hash, id, size)

	if err != nil {
		t.Errorf("%s", err)

	}

	t.Errorf("%s", ret)

}
