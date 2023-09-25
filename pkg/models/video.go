package models

import (
	"fmt"
	"reflect"
	"strings"
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

type VideoQueryParams struct {
	ChannelID          string
	ID                 string
	Include            string
	Lang               enums.LangsType
	Limit              int
	MaxUpcomingHours   int
	MentionedChannelID string
	Offset             int
	Order              enums.OrderType
	Org                string
	Paginated          string
	Sort               enums.VideoSortBy_Type
	Status             enums.VideoStatusType
	Topic              string
	Type               enums.VideoType_Type
}

func (v VideoQueryParams) GetQueryString() string {
	var sb strings.Builder

	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)

		name := strings.ToLower(field.Name)

		switch field.Type.Kind() {
		case reflect.String:
			if value.String() == "" {
				continue
			}

			sb.WriteString(fmt.Sprintf("%s=%s&", name, value.String()))
		case reflect.Int:
			if value.Int() == 0 {
				continue
			}

			sb.WriteString(fmt.Sprintf("%s=%d&", name, value.Int()))
		}
	}

	if sb.Len() > 0 {
		return "?" + sb.String()[:sb.Len()-1]
	}
	return ""
}
