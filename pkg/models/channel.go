package models

import (
	"fmt"
	"reflect"
	"strings"
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

type ChannelQueryParams struct {
	Type   enums.ChannelType_Type
	Offset int
	Limit  int
	Org    string
	Lang   enums.LangsType
	Sort   enums.ChannelSortBy_Type
	Order  enums.OrderType
}

func (c ChannelQueryParams) GetQueryString() string {
	var sb strings.Builder

	val := reflect.ValueOf(c)
	typ := reflect.TypeOf(c)

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
