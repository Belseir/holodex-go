package models

import (
	"time"

	"github.com/belseir/holodex-go/pkg/enums"
)

type Video struct {
	ID             string                `json:"id"`
	Title          string                `json:"title"`
	Type           enums.VideoType_Type  `json:"type"`
	TopicID        string                `json:"topic_id,omitempty"`
	PublishedAt    time.Time             `json:"published_at,omitempty"`
	AvailableAt    time.Time             `json:"available_at"`
	Duration       int                   `json:"duration"`
	Status         enums.VideoStatusType `json:"status"`
	StartScheduled time.Time             `json:"start_scheduled,omitempty"`
	StartActual    time.Time             `json:"start_actual,omitempty"`
	EndActual      time.Time             `json:"end_actual,omitempty"`
	LiveViewers    int                   `json:"live_viewers,omitempty"`
	Description    string                `json:"description,omitempty"`
	Songcount      int                   `json:"songcount"`
	ChannelID      string                `json:"channel_id"`
}

type VideoWithChannel struct {
	Video
	Channel ChannelMin `json:"channel"`
}

type VideoFull struct {
	Video
	Clips      []VideoWithChannel  `json:"clips,omitempty"`
	Sources    []VideoWithChannel  `json:"sources,omitempty"`
	Refers     []VideoWithChannel  `json:"refers,omitempty"`
	Simulcasts []VideoWithChannel  `json:"simulcasts,omitempty"`
	Mentions   []ChannelMinWithOrg `json:"mentions,omitempty"`
	Songs      int                 `json:"songs,omitempty"`
}
