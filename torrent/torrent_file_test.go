package torrent

import (
	"testing"
)

func TestNewTorrentFromString(t *testing.T) {
	type testCase struct {
		data     string
		expected File
	}

	testCases := []testCase{
		{
			`d8:announce43:udp://tracker.coppersurfer.tk:6969/announce10:created by13:uTorrent/187013:creation datei1462355939e8:encoding5:UTF-84:infod6:lengthi124234e4:name9:puppy.jpg12:piece lengthi16384e6:pieces160:T�k�/�_(�S\u0011h%���+]q\'B\u0018�٠:����p"�j����1-g"\u0018�s(\u001b\u000f���V��=�h�m\u0017a�nF�2���N\r�ǩ�_�\u001e"2���\'�wO���-;\u0004ע\u0017�ؑ��L&����0\u001f�D_9��\t\\��O�h,n\u001a5g�(��仑,�\\߰�%��U��\u0019��C\u0007>��df��ee`,
			File{
				"udp://tracker.coppersurfer.tk:6969/announce",
				"uTorrent/1870",
				1462355939,
				"UTF-8",
				info{
					124234,
					"puppy.jpg",
					16384,
					`T�k�/�_(�S\\u0011h%���+]q\'B\\u0018�٠:����p\\"�j����1-g\\"\\u0018�s(\\u001b\\u000f���V��=�h�m\\u0017a�nF�2���N\\r�ǩ�_�\\u001e\\"2���\'�wO���-;\\u0004ע\\u0017�ؑ��L&����0\\u001f�D_9��\\t\\\\��O�h,n\\u001a5g�(��仑,�\\\\߰�%��U��\\u0019��C\\u0007>��df��`,
				},
			},
		},
	}

	for _, test := range testCases {
		var torrent File
		torrent = NewTorrentFromString(test.data)

		if torrent.Encoding != test.expected.Encoding {
			t.Errorf("error: Encoding %s is not equal to %s", torrent.Encoding, test.expected.Encoding)
		}
		if torrent.Announce != test.expected.Announce {
			t.Errorf("error: Announce %s is not equal to %s", torrent.Announce, test.expected.Announce)
		}
		if torrent.CreatedBy != test.expected.CreatedBy {
			t.Errorf("error: CreatedBy %s is not equal to %s", torrent.CreatedBy, test.expected.CreatedBy)
		}
		if torrent.CreationDate != test.expected.CreationDate {
			t.Errorf("error: CreationDate %v is not equal to %v", torrent.CreationDate, test.expected.CreationDate)
		}
		if torrent.Info.Length != test.expected.Info.Length {
			t.Errorf("error: Info.Length %v is not equal to %v", torrent.Info.Length, test.expected.Info.Length)
		}
		if torrent.Info.Name != torrent.Info.Name {
			t.Errorf("error: Info.Length %s is not equal to %s", torrent.Info.Name, test.expected.Info.Name)
		}
		if torrent.Info.Pieces != torrent.Info.Pieces {
			t.Errorf("error: Info.Pieces %s is not equal to %s", torrent.Info.Pieces, test.expected.Info.Pieces)
		}
		if torrent.Info.PieceLength != torrent.Info.PieceLength {
			t.Errorf("error: Info.PieceLength %v is not equal to %v", torrent.Info.PieceLength, test.expected.Info.PieceLength)
		}
	}

}
