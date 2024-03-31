package spacial_queries

import (
	"fmt"
	"github.com/mmcloughlin/geohash"
)

type Point[T any] struct {
	lat float64
	lng float64
	v   T
}

func (r *Point[T]) Encode() uint64 {
	return geohash.EncodeInt(r.lat, r.lng)
}

func DecodeToPoint[T any](hash uint64, v T) Point[T] {
	lat, lng := geohash.DecodeInt(hash)

	return Point[T]{
		lat: lat,
		lng: lng,
		v:   v,
	}
}

func (r *Point[T]) String() string {
	return fmt.Sprintf("%.3f, %.3f: %v", r.lat, r.lng, r.v)
}
