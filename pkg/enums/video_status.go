package enums

import "fmt"

var VideoStatus = newVideoStatus()

type VideoStatusType string

func newVideoStatus() *videoStatus {
	return &videoStatus{
		NEW:      "new",
		UPCOMING: "upcoming",
		LIVE:     "live",
		PAST:     "past",
		MISSING:  "missing",
	}
}

type videoStatus struct {
	NEW      VideoStatusType
	UPCOMING VideoStatusType
	LIVE     VideoStatusType
	PAST     VideoStatusType
	MISSING  VideoStatusType
}

func (_ *videoStatus) IsValid(vs VideoStatusType) (bool, error) {
	switch vs {
	case
		VideoStatus.NEW,
		VideoStatus.UPCOMING,
		VideoStatus.LIVE,
		VideoStatus.PAST,
		VideoStatus.MISSING:
		return true, nil
	default:
		return false, fmt.Errorf("Invalid VideoStatus (%s) accessed", vs)
	}
}
