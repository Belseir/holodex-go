package models

import (
	"time"

	"github.com/belseir/holodex-go/pkg/enums"
)

type Channel struct {
	ID              string                 `json:"id"`
	Name            string                 `json:"name"`
	EnglishName     string                 `json:"english_name,omitempty"`
	Type            enums.ChannelType_Type `json:"type"`
	Org             string                 `json:"org,omitempty"`
	Suborg          string                 `json:"suborg,omitempty"`
	Photo           string                 `json:"photo,omitempty"`
	Banner          string                 `json:"banner,omitempty"`
	Twitter         string                 `json:"twitter,omitempty"`
	VideoCount      string                 `json:"video_count,omitempty"`
	SubscriberCount string                 `json:"subscriber_count,omitempty"`
	ViewCount       string                 `json:"view_count,omitempty"`
	ClipCount       string                 `json:"clip_count,omitempty"`
	Lang            enums.LangsType        `json:"lang,omitempty"`
	PublishedAt     time.Time              `json:"published_at"`
	Inactive        bool                   `json:"inactive"`
	Description     string                 `json:"description"`
}

type ChannelMin struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	EnglishName string                 `json:"english_name,omitempty"`
	Type        enums.ChannelType_Type `json:"type"`
	Photo       string                 `json:"photo"`
}

type ChannelMinWithOrg struct {
	ChannelMin
	Org string `json:"org,omitempty"`
}

type ChannelWithGroup struct {
	Channel
	Group string `json:"group,omitempty"`
}
