package tracker

import (
	"testing"
	unsafe "unsafe"
)

func TestTrackerProtocolStructs(t *testing.T) {

	var creq ConnectRequest

	if unsafe.Sizeof(creq) != 16 {
		t.Errorf("error: ConnectionRequest has not the proper size %v => %v", 16, unsafe.Sizeof(creq))
	}

	var cres ConnectResponse

	if unsafe.Sizeof(cres) != 16 {
		t.Errorf("error: ConnectResponse has not the proper size %v => %v", 16, unsafe.Sizeof(cres))
	}

	var areq AnnounceRequest

	if unsafe.Sizeof(areq) != 104 {
		t.Errorf("error: AnnounceRequest has not the proper size %v => %v", 104, unsafe.Sizeof(areq))
	}
}
