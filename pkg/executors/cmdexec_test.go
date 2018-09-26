package executors

import (
	"reflect"
	"testing"
)

func TestNewCmdExecutor(t *testing.T) {
	cases := []struct {
		input string
		want  *cmdExecutor
	}{
		{"ls", &cmdExecutor{"ls"}},
		{"slkdfj", &cmdExecutor{}},
	}
	for _, c := range cases {
		got := NewCmdExecutor(c.input)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("NewCmdExecutor(%v) == %v, wanted %v", c.input, got, c.want)
		}
	}
}
