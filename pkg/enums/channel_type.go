package enums

import "fmt"

var ChannelType = newChannelType()

type ChannelType_Type string

func newChannelType() *channelType {
	return &channelType{
		VTUBER: "vtuber",
		SUBBER: "subber",
	}
}

type channelType struct {
	VTUBER ChannelType_Type
	SUBBER ChannelType_Type
}

func (_ *channelType) IsValid(ct ChannelType_Type) (bool, error) {
	if ct != ChannelType.VTUBER && ct != ChannelType.SUBBER {
		return false, fmt.Errorf("Invalid ChannelType (%s) accessed. Expected 'vtuber' or 'subber'", ct)
	}

	return true, nil
}
