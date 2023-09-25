package builders

import (
	"fmt"

	"github.com/belseir/holodex-go/pkg/enums"
	"github.com/belseir/holodex-go/pkg/models"
)

type IVideoQueryBuilder interface {
	SetChannelID(id string) *videoQueryBuilder
	SetID(id string) *videoQueryBuilder
	SetInclude(i string) *videoQueryBuilder
	SetLang(l enums.LangsType) *videoQueryBuilder
	SetLimit(l int) *videoQueryBuilder
	SetMaxUpcomingHours(h int) *videoQueryBuilder
	SetMentionedChannelID(id string) *videoQueryBuilder
	SetOffset(o int) *videoQueryBuilder
	SetOrder(o enums.OrderType) *videoQueryBuilder
	SetOrg(o string) *videoQueryBuilder
	SetPaginated(p string) *videoQueryBuilder
	SetSort(s enums.VideoSortBy_Type) *videoQueryBuilder
	SetStatus(s enums.VideoStatusType) *videoQueryBuilder
	SetTopic(t string) *videoQueryBuilder
	SetType(t enums.VideoType_Type) *videoQueryBuilder
	validate() error
	Build() (*models.VideoQueryParams, error)
}

type videoQueryBuilder struct {
	channelId          string
	id                 string
	include            string
	lang               enums.LangsType
	limit              int
	maxUpcomingHours   int
	mentionedChannelId string
	offset             int
	order              enums.OrderType
	org                string
	paginated          string
	sort               enums.VideoSortBy_Type
	status             enums.VideoStatusType
	topic              string
	_type              enums.VideoType_Type
}

func NewVideoQueryBuilder() *videoQueryBuilder {
	return &videoQueryBuilder{}
}

func (b *videoQueryBuilder) SetChannelID(id string) *videoQueryBuilder {
	b.channelId = id
	return b
}

func (b *videoQueryBuilder) SetID(id string) *videoQueryBuilder {
	b.id = id
	return b
}

func (b *videoQueryBuilder) SetInclude(i string) *videoQueryBuilder {
	b.include = i
	return b
}

func (b *videoQueryBuilder) SetLang(l enums.LangsType) *videoQueryBuilder {
	b.lang = l
	return b
}

func (b *videoQueryBuilder) SetLimit(l int) *videoQueryBuilder {
	b.limit = l
	return b
}

func (b *videoQueryBuilder) SetMaxUpcomingHours(h int) *videoQueryBuilder {
	b.maxUpcomingHours = h
	return b
}

func (b *videoQueryBuilder) SetMentionedChannelID(id string) *videoQueryBuilder {
	b.mentionedChannelId = id
	return b
}

func (b *videoQueryBuilder) SetOffset(o int) *videoQueryBuilder {
	b.offset = o
	return b
}

func (b *videoQueryBuilder) SetOrder(o enums.OrderType) *videoQueryBuilder {
	b.order = o
	return b
}

func (b *videoQueryBuilder) SetOrg(o string) *videoQueryBuilder {
	b.org = o
	return b
}

func (b *videoQueryBuilder) SetPaginated(p string) *videoQueryBuilder {
	b.paginated = p
	return b
}

func (b *videoQueryBuilder) SetSort(s enums.VideoSortBy_Type) *videoQueryBuilder {
	b.sort = s
	return b
}

func (b *videoQueryBuilder) SetStatus(s enums.VideoStatusType) *videoQueryBuilder {
	b.status = s
	return b
}

func (b *videoQueryBuilder) SetTopic(t string) *videoQueryBuilder {
	b.topic = t
	return b
}

func (b *videoQueryBuilder) SetType(t enums.VideoType_Type) *videoQueryBuilder {
	b._type = t
	return b
}

func (b *videoQueryBuilder) validate() error {

	if b.limit > MAX_LIMIT {
		return fmt.Errorf("Invalid Limit (%d) in videoQueryBuilder.setLimit(...). Expected to be <= 50", b.limit)
	}

	if b.lang != "" {
		if ok, err := enums.Langs.IsValid(b.lang); !ok {
			return err
		}
	}

	if b.order != "" {
		if ok, err := enums.Order.IsValid(b.order); !ok {
			return err
		}
	}

	if b.sort != "" {
		if ok, err := enums.VideoSortBy.IsValid(b.sort); !ok {
			return err
		}
	}

	if b.status != "" {
		if ok, err := enums.VideoStatus.IsValid(b.status); !ok {
			return err
		}
	}

	if b._type != "" {
		if ok, err := enums.VideoType.IsValid(b._type); !ok {
			return err
		}
	}

	return nil
}

func (b *videoQueryBuilder) Build() (*models.VideoQueryParams, error) {
	if err := b.validate(); err != nil {
		return nil, err
	}

	return &models.VideoQueryParams{
		ChannelID:          b.channelId,
		ID:                 b.id,
		Include:            b.include,
		Lang:               b.lang,
		Limit:              b.limit,
		MaxUpcomingHours:   b.maxUpcomingHours,
		MentionedChannelID: b.mentionedChannelId,
		Offset:             b.offset,
		Order:              b.order,
		Org:                b.org,
		Paginated:          b.paginated,
		Sort:               b.sort,
		Status:             b.status,
		Topic:              b.topic,
		Type:               b._type,
	}, nil
}
