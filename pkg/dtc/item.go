package dtc

import (
	"encoding/json"
	"path/filepath"
	"strconv"
	"strings"
)

type Item struct {
	FileName   string
	SubPath    string
	Crop       string
	Filters    []string
	ForcedRate string
	width      int
	height     int
}

func NewItem(filename string, width, height int) *Item {
	i := new(Item)
	i.FileName = filename
	i.width = width
	i.height = height
	i.Crop = "0:0:0:0"
	return i
}

func NewItemFromPath(path string, inboxPath string) *Item {
	i := new(Item)
	i.FileName = filepath.Base(path)
	relativePath, _ := filepath.Rel(inboxPath, path)
	i.SubPath, _ = filepath.Split(relativePath)
	return i
}

func (i *Item) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}

func (i *Item) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &i); err != nil {
		return err
	}
	return nil
}

func (i *Item) SetCrop(value string) {
	if validCrop(value, i.width, i.height) {
		i.Crop = value
	}
}

func validCrop(crop string, width int, height int) bool {
	if width <= 0 || height <= 0 {
		// what is cropped may never crop
		return false
	}
	if strings.Count(crop, ":") != 3 {
		// valid format n:n:n:n
		return false
	}

	values := strings.Split(crop, ":")
	if len(values) != 4 {
		// we never get here because a string with 3 colons always splits into
		// 4 strings.  Probably.  But here's the check anyway.
		return false
	}
	var intvalues []int
	for _, v := range values {
		iv, err := strconv.Atoi(v)
		if err != nil {
			// check that each substring is an int
			return false
		} else if iv < 0 || iv > 10000 {
			return false
		}
		intvalues = append(intvalues, iv)
	}
	if intvalues[0]+intvalues[1] >= width || intvalues[2]+intvalues[3] >= height {
		// also catches the case when 1 of them is greater than width
		return false
	}
	// if we've survived the guantlet, then it's a valid crop value
	return true
}
