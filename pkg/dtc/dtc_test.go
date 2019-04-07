package dtc

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerateTranscodeVideoCommand(t *testing.T) {
	cases := []struct {
		workspace  string
		crop       string
		forcedRate string
		filters    []string
		inboxpath  string
		outboxpath string
		fullpath   string
		want       []string
	}{
		{
			"/Volumes/transcode",
			"0:0:0:0",
			"",
			[]string{},
			"inbox/",
			"outbox/",
			"/Volumes/transcode/inbox/Airplane! (1980)/Airplane! (1980).mkv",
			[]string{
				"transcode-video",
				"--no-log",
				"--crop", "0:0:0:0",
				"--output", "/Volumes/transcode/outbox/Airplane! (1980)/Airplane! (1980).mkv",
				"/Volumes/transcode/inbox/Airplane! (1980)/Airplane! (1980).mkv",
			},
		},
		{
			"/data",
			"1:2:3:4",
			"23.976",
			[]string{"detelecine"},
			"inbox/",
			"outbox/",
			"/data/inbox/Frasier/Season 01/Frasier - s01e01.mkv",
			[]string{
				"transcode-video",
				"--no-log",
				"--crop", "1:2:3:4",
				"--force-rate", "23.976",
				"--filter", "detelecine",
				"--output", "/data/outbox/Frasier/Season 01/Frasier - s01e01.mkv",
				"/data/inbox/Frasier/Season 01/Frasier - s01e01.mkv",
			},
		},
	}
	for _, c := range cases {
		j := new(Job)
		j.InboxPath = c.inboxpath
		j.OutboxPath = c.outboxpath
		j.Item = NewItemFromPath(c.fullpath, filepath.Join(c.workspace,
			j.InboxPath))
		j.Item.Crop = c.crop
		j.Item.ForcedRate = c.forcedRate
		j.Item.Filters = c.filters
		got := generateTranscodeVideoCommand(c.workspace, j)
		if strings.Join(got, " ") != strings.Join(c.want, " ") {
			t.Errorf("generateTranscodeVideoCommand(%v, %v) == %v, wanted %v",
				c.workspace, j, strings.Join(got, " "),
				strings.Join(c.want, " "))
		}
	}
}
