package executors

import (
	"testing"

	"github.com/crerwin/distributedtranscoding/pkg/mock"
)

func TestDetectCrop(t *testing.T) {
	cases := []struct {
		args []string
		want string
	}{
		{[]string{}, "detect-crop"},
	}
	e := videoTranscodingExecutor{executor: mock.NewCmdExecutor("detect-crop")}
	for _, c := range cases {
		got := e.DetectCrop(c.args...)
		if got != c.want {
			t.Errorf("videoTranscodingExecutor.DetectCrop(%v) == %v, wanted %v", c.args, got, c.want)
		}
	}
}

func TestTranscodeVideo(t *testing.T) {
	cases := []struct {
		args []string
		want string
	}{
		{[]string{}, "transcode-video"},
		{[]string{"title00.mkv"}, "transcode-video title00.mkv"},
	}
	e := videoTranscodingExecutor{executor: mock.NewCmdExecutor("transcode-video")}
	for _, c := range cases {
		got := e.TranscodeVideo(c.args...)
		if got != c.want {
			t.Errorf("videoTranscodingExecutor.TranscodeVideo(%v) == %v, wanted %v", c.args, got, c.want)
		}
	}
}
