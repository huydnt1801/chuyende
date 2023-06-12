package trip

import (
	"github.com/huydnt1801/chuyende/internal/ent"
	"github.com/mitchellh/mapstructure"
)

func DecodeTrip(input *ent.Trip) (*Trip, error) {
	u := &Trip{}
	err := mapstructure.Decode(input, u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func MustDecodeTrip(input *ent.Trip) *Trip {
	out, err := DecodeTrip(input)
	if err != nil {
		panic(err)
	}
	return out
}

func DecodeManyTrip(input []*ent.Trip) ([]*Trip, error) {
	ret := make([]*Trip, len(input))
	for i, m := range input {
		r, err := DecodeTrip(m)
		if err != nil {
			return nil, err
		}
		ret[i] = r
	}
	return ret, nil
}
