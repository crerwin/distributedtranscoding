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
		{[]string{}, "detect-crop mock output"},
	}
	e := videoTranscodingExecutor{executor: mock.NewCmdExecutor("detect-crop")}
	for _, c := range cases {
		got := e.DetectCrop(c.args...)
		if got != c.want {
			t.Errorf("videoTranscodingExecutor.DetectCrop(%v) == %v, wanted %v", c.args, got, c.want)
		}
	}
}
