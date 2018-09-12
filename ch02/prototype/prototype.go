package prototype

import (
	"errors"
	"fmt"
)

type ShirtCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}

const (
	White = iota + 1
	Black
	Blue
)

func GetShirtsCloner() ShirtCloner {
	return new(ShirtsCache)
}

type ShirtsCache struct{}

func (i *ShirtsCache) GetClone(s int) (ItemInfoGetter, error) {
	var newShirt Shirt
	switch s {
	case White:
		newShirt = *whitePrototype
	case Black:
		newShirt = *blackPrototype
	case Blue:
		newShirt = *bluePrototype
	default:
		return nil, errors.New(fmt.Sprintf("Unknown color code: %d", s))
	}
	return &newShirt, nil

}

type ItemInfoGetter interface {
	GetInfo() string
}

type ShirtColor byte

type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt with SKU '%s' and Color id %d that costs %f\n", s.SKU, s.Color, s.Price)
}

var whitePrototype *Shirt = &Shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: White,
}

var blackPrototype *Shirt = &Shirt{
	Price: 16.00,
	SKU:   "empty",
	Color: Black,
}

var bluePrototype *Shirt = &Shirt{
	Price: 17.00,
	SKU:   "empty",
	Color: Blue,
}

func (i *Shirt) GetPrice() float32 {
	return i.Price
}
