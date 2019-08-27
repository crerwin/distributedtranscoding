package dtc

import "testing"

func TestCreateJobName(t *testing.T) {
	cases := []struct {
		prefix   string
		fileName string
		want     string
	}{
		{"dtc-", "Frasier - s01e01.mkv", "dtc-frasiers01e01"},
		{"dtc-", "Airplane! (1980)/Airplane! (1980).mkv", "dtc-airplane1980airplane1980"},
		{"dtc-", "Master and Commander - The Far Side of the World (2003).mkv",
			"dtc-masterandcommanderthefarsideoftheworld2003"},
		{"dtc-", "The Naked Gun-  From The Files Of Police Squad! (1988)/The Naked Gun-  From The Files Of Police Squad! (1988).mkv",
			"dtc-esofpolicesquad1988thenakedgunfromthefilesofpolicesquad1988"},
		{"longprefix",
			"Night Of The Day Of The Dawn Of The Son Of The Bride Of The Return Of The Revenge Of The Terror Of The Attack Of The Evil, Mutant, Hellbound, Flesh-Eating, Crawling, Alien, Zombified, Subhumanoid Living Dead — Part 5",
			"longprefixatingcrawlingalienzombifiedsubhumanoidlivingdeadpart5"},
	}
	for _, c := range cases {
		got := createJobName(c.prefix, c.fileName)
		if got != c.want {
			t.Errorf("createJobName(%v, %v) == %v, want %v", c.prefix,
				c.fileName, got, c.want)
		}
	}
}
