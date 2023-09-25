package enums

import "fmt"

var VideoType = newVideoType()

type VideoType_Type string

func newVideoType() *videoType {
	return &videoType{
		STREAM: "stream",
		CLIP:   "clip",
	}
}

type videoType struct {
	STREAM VideoType_Type
	CLIP   VideoType_Type
}

func (v *videoType) IsValid(t VideoType_Type) (bool, error) {
	if t != VideoType.STREAM && t != VideoType.CLIP {
		return false, fmt.Errorf("Invalid VideoType (%s) accessed. Expected 'stream' or 'clip'", t)
	}

	return true, nil
}
