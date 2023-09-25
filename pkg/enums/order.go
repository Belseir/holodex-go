package enums

import "fmt"

var Order = newOrder()

type OrderType string

func newOrder() *order {
	return &order{
		ASC:  "asc",
		DESC: "desc",
	}
}

type order struct {
	ASC  OrderType
	DESC OrderType
}

func (o *order) IsValid(ot OrderType) (bool, error) {
	if ot != Order.ASC && ot != Order.DESC {
		return false, fmt.Errorf("Invalid Order (%s) accessed. Expected 'asc' or 'desc'", ot)
	}

	return true, nil
}
