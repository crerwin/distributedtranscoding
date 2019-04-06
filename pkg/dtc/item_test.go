package dtc

import "testing"

func TestValidCrop(t *testing.T) {
	cases := []struct {
		input  string
		width  int
		height int
		want   bool
	}{
		{"hello", 10, 10, false},
		{"0:0:0", 10, 10, false},
		{"0:0::0", 10, 10, false},
		{"0:0:4:5", 10, 10, true},
		{"f:0:0:0", 10, 10, false},
		{"0:-1:0:0", 10, 10, false},
		{"0:14321:0:0", 10, 10, false},
		{"0:0:0:0", 0, 0, false},
		{"10:0:0:0", 5, 5, false},
		{"6:4:0:0", 10, 10, false},
	}
	for _, c := range cases {
		got := validCrop(c.input, c.width, c.height)
		if got != c.want {
			t.Errorf("validCrop(%v, %v, %v) == %v, wanted %v",
				c.input, c.width, c.height, got, c.want,
			)
		}
	}
}

func TestNewItemFromPath(t *testing.T) {
	cases := []struct {
		path           string
		inboxPath      string
		wantedFileName string
		wantedSubPath  string
	}{
		{
			"/Volumes/transcode/inbox/frasier/s1d1/Frasier - s01e01.mkv",
			"/Volumes/transcode/inbox/",
			"Frasier - s01e01.mkv",
			"frasier/s1d1/",
		}, {
			"/data/inbox/Airplane! (1980)/Airplane! (1980).mkv",
			"/data/inbox/",
			"Airplane! (1980).mkv",
			"Airplane! (1980)/",
		},
	}
	for _, c := range cases {
		gotItem := NewItemFromPath(c.path, c.inboxPath)
		if gotItem.FileName != c.wantedFileName {
			t.Errorf("NewItemFromPath(%v, %v).FileName == %v, wanted %v",
				c.path, c.inboxPath, gotItem.FileName, c.wantedFileName,
			)
		}
		if gotItem.SubPath != c.wantedSubPath {
			t.Errorf("NewItemFromPath(%v, %v).SubPath == %v, wanted %v",
				c.path, c.inboxPath, gotItem.SubPath, c.wantedSubPath,
			)
		}
	}
}
