package builders

import (
	"fmt"

	"github.com/belseir/holodex-go/pkg/enums"
	"github.com/belseir/holodex-go/pkg/models"
)

type IChannelQueryBuilder interface {
	SetType(t enums.ChannelType_Type) *channelQueryBuilder
	SetOffSet(o int) *channelQueryBuilder
	SetLimit(l int) *channelQueryBuilder
	SetOrg(o string) *channelQueryBuilder
	SetLang(l enums.LangsType) *channelQueryBuilder
	SetSort(s enums.ChannelSortBy_Type) *channelQueryBuilder
	SetOrder(o enums.OrderType) *channelQueryBuilder
	validate() error
	Build() (*models.ChannelQueryParams, error)
}

type channelQueryBuilder struct {
	_type  enums.ChannelType_Type
	offset int
	limit  int
	org    string
	lang   enums.LangsType
	sort   enums.ChannelSortBy_Type
	order  enums.OrderType
}

const MAX_LIMIT = 50

func NewChannelQueryBuilder() *channelQueryBuilder {
	return &channelQueryBuilder{}
}

func (b *channelQueryBuilder) SetType(t enums.ChannelType_Type) *channelQueryBuilder {
	b._type = t
	return b
}

func (b *channelQueryBuilder) SetOffSet(o int) *channelQueryBuilder {
	b.offset = o
	return b
}

func (b *channelQueryBuilder) SetLimit(l int) *channelQueryBuilder {
	b.limit = l
	return b
}

func (b *channelQueryBuilder) SetOrg(o string) *channelQueryBuilder {
	b.org = o
	return b
}

func (b *channelQueryBuilder) SetLang(l enums.LangsType) *channelQueryBuilder {
	b.lang = l
	return b
}

func (b *channelQueryBuilder) SetSort(s enums.ChannelSortBy_Type) *channelQueryBuilder {
	b.sort = s
	return b
}

func (b *channelQueryBuilder) SetOrder(o enums.OrderType) *channelQueryBuilder {
	b.order = o
	return b
}

func (b *channelQueryBuilder) validate() error {
	if b._type != "" {
		if ok, err := enums.ChannelType.IsValid(b._type); !ok {
			return err
		}
	}

	if b.limit > 50 {
		return fmt.Errorf("Invalid Limit (%d) in ChannelQueryBuilder.setLimit(...). Expected to be <= 50", b.limit)
	}

	if b.lang != "" {
		if ok, err := enums.Langs.IsValid(b.lang); !ok {
			return err
		}
	}

	if b.sort != "" {
		if ok, err := enums.ChannelSortBy.IsValid(b.sort); !ok {
			return err
		}
	}

	if b.order != "" {
		if ok, err := enums.Order.IsValid(b.order); !ok {
			return err
		}
	}

	return nil
}

func (b *channelQueryBuilder) Build() (*models.ChannelQueryParams, error) {
	if err := b.validate(); err != nil {
		return nil, err
	}

	return &models.ChannelQueryParams{
		Type:   b._type,
		Offset: b.offset,
		Limit:  b.limit,
		Org:    b.org,
		Lang:   b.lang,
		Sort:   b.sort,
		Order:  b.order,
	}, nil
}
