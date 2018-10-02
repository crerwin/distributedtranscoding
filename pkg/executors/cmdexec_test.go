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

func TestExecute(t *testing.T) {
	cases := []struct {
		cmd  string
		args []string
		want string
	}{
		{"echo", []string{}, ""},
		{"echo", []string{"test"}, "test"},
		{"echo", []string{"test1", "test2"}, "test1 test2"},
	}
	for _, c := range cases {
		ce := NewCmdExecutor(c.cmd)
		got, _ := ce.Execute(c.args...)
		if got != c.want {
			t.Errorf("Execute() == %v, wanted %v", got, c.want)
		}
	}
}
