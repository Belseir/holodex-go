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

func (_ *order) IsValid(o OrderType) (bool, error) {
	if o != Order.ASC && o != Order.DESC {
		return false, fmt.Errorf("Invalid Order (%s) accessed. Expected 'asc' or 'desc'", o)
	}

	return true, nil
}
