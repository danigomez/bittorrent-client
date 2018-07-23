package torrent

import (
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
