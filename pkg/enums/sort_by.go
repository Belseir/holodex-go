package enums

import "fmt"

var ChannelSortBy = newChannelSortBy()
var VideoSortBy = newVideoSortBy()

type ChannelSortBy_Type string

func newChannelSortBy() *channelSortBy {
	return &channelSortBy{
		ID:               "id",
		NAME:             "name",
		ENGLISH_NAME:     "englis_hname",
		TYPE:             "type",
		ORG:              "org",
		SUBORG:           "suborg",
		PHOTO:            "photo",
		BANNER:           "banner",
		TWITTER:          "twitter",
		VIDEO_COUNT:      "video_count",
		SUBSCRIBER_COUNT: "subscriber_count",
		VIEW_COUNT:       "view_count",
		CLIP_COUNT:       "clip_count",
		LANG:             "lang",
		PUBLISHED_AT:     "published_at",
		INACTIVE:         "inactive",
		DESCRIPTION:      "description",
		GROUP:            "group",
	}
}

type channelSortBy struct {
	ID               ChannelSortBy_Type
	NAME             ChannelSortBy_Type
	ENGLISH_NAME     ChannelSortBy_Type
	TYPE             ChannelSortBy_Type
	ORG              ChannelSortBy_Type
	SUBORG           ChannelSortBy_Type
	PHOTO            ChannelSortBy_Type
	BANNER           ChannelSortBy_Type
	TWITTER          ChannelSortBy_Type
	VIDEO_COUNT      ChannelSortBy_Type
	SUBSCRIBER_COUNT ChannelSortBy_Type
	VIEW_COUNT       ChannelSortBy_Type
	CLIP_COUNT       ChannelSortBy_Type
	LANG             ChannelSortBy_Type
	PUBLISHED_AT     ChannelSortBy_Type
	INACTIVE         ChannelSortBy_Type
	DESCRIPTION      ChannelSortBy_Type
	GROUP            ChannelSortBy_Type
}

type VideoSortBy_Type string

func newVideoSortBy() *videoSortBy {
	return &videoSortBy{
		ID:              "id",
		TITLE:           "title",
		TYPE:            "type",
		TOPIC_ID:        "topic_id",
		PUBLISHED_AT:    "published_at",
		AVAILABLE_AT:    "available_at",
		DURATION:        "duration",
		STATUS:          "status",
		START_SCHEDULED: "start_scheduled",
		START_ACTUAL:    "start_actual",
		END_ACTUAL:      "end_actual",
		LIVE_VIEWERS:    "live_viewers",
		DESCRIPTION:     "description",
		SONGCOUNT:       "songcount",
		SONGS:           "songs",
	}
}

type videoSortBy struct {
	ID              VideoSortBy_Type
	TITLE           VideoSortBy_Type
	TYPE            VideoSortBy_Type
	TOPIC_ID        VideoSortBy_Type
	PUBLISHED_AT    VideoSortBy_Type
	AVAILABLE_AT    VideoSortBy_Type
	DURATION        VideoSortBy_Type
	STATUS          VideoSortBy_Type
	START_SCHEDULED VideoSortBy_Type
	START_ACTUAL    VideoSortBy_Type
	END_ACTUAL      VideoSortBy_Type
	LIVE_VIEWERS    VideoSortBy_Type
	DESCRIPTION     VideoSortBy_Type
	SONGCOUNT       VideoSortBy_Type
	SONGS           VideoSortBy_Type
}

func (c *channelSortBy) IsValid(s ChannelSortBy_Type) (bool, error) {
	switch s {
	case
		ChannelSortBy.ID,
		ChannelSortBy.NAME,
		ChannelSortBy.ENGLISH_NAME,
		ChannelSortBy.TYPE,
		ChannelSortBy.ORG,
		ChannelSortBy.SUBORG,
		ChannelSortBy.PHOTO,
		ChannelSortBy.BANNER,
		ChannelSortBy.TWITTER,
		ChannelSortBy.VIDEO_COUNT,
		ChannelSortBy.SUBSCRIBER_COUNT,
		ChannelSortBy.VIEW_COUNT,
		ChannelSortBy.CLIP_COUNT,
		ChannelSortBy.LANG,
		ChannelSortBy.PUBLISHED_AT,
		ChannelSortBy.INACTIVE,
		ChannelSortBy.DESCRIPTION,
		ChannelSortBy.GROUP:
		return true, nil
	default:
		return false, fmt.Errorf("Invalid ChannelSortBy (%s) accessed", s)
	}
}

func (v *videoSortBy) IsValid(s VideoSortBy_Type) (bool, error) {
	switch s {
	case
		VideoSortBy.ID,
		VideoSortBy.TITLE,
		VideoSortBy.TYPE,
		VideoSortBy.TOPIC_ID,
		VideoSortBy.PUBLISHED_AT,
		VideoSortBy.AVAILABLE_AT,
		VideoSortBy.DURATION,
		VideoSortBy.STATUS,
		VideoSortBy.START_SCHEDULED,
		VideoSortBy.START_ACTUAL,
		VideoSortBy.END_ACTUAL,
		VideoSortBy.LIVE_VIEWERS,
		VideoSortBy.DESCRIPTION,
		VideoSortBy.SONGCOUNT,
		VideoSortBy.SONGS:
		return true, nil
	default:
		return false, fmt.Errorf("Invalid VideoSortBy (%s) accessed", s)
	}
}
