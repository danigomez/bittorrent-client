package torrent

import "testing"

func TestGetId(t *testing.T) {

	testNames := []string {
		"-BBV01-",
		"-BBV010-",
		"-BBV0100-",
		"-BBV01000-",
		"-BBV000011-",
		"-BBV0100000000000000000000000000000000000000000000000000000-",
	}

	for i, name := range testNames {

		if i == 5 {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("The code did not panic")
				}
			}()
		}

		id := GetId(name)

		if len(id) != 20 {
			t.Errorf("error: Id is not generated with the correct len it should be 20, and is %v", len(id))
		}


	}

}
