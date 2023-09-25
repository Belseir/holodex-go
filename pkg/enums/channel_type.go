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

func (c *channelType) IsValid(t ChannelType_Type) (bool, error) {
	if t != ChannelType.VTUBER && t != ChannelType.SUBBER {
		return false, fmt.Errorf("Invalid ChannelType (%s) accessed. Expected 'vtuber' or 'subber'", t)
	}

	return true, nil
}
