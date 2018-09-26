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
			t.Errorf("validCrop(%v, %v, %v) == %v, wanted %v", c.input, c.width, c.height, got, c.want)
		}
	}
}
