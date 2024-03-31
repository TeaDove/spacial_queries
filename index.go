package spacial_queries

import (
	"fmt"
	"github.com/mmcloughlin/geohash"
	"github.com/tidwall/btree"
	"strings"
)

type GeohashBTree[T any] struct {
	index btree.Map[uint64, T]
}

func NewGeohashBTree[T any]() *GeohashBTree[T] {
	tree := GeohashBTree[T]{index: *btree.NewMap[uint64, T](4)}
	return &tree
}

type IterFunc[T any] func(point Point[T]) bool

func (r *GeohashBTree[T]) String() string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("Len: %d\n", r.index.Len()))

	r.Iter(func(point Point[T]) bool {
		builder.WriteString(fmt.Sprintf("%s\n", point.String()))
		return true
	})

	return builder.String()
}

func (r *GeohashBTree[T]) Set(point *Point[T]) {
	r.index.Set(point.Encode(), point.v)
}

func (r *GeohashBTree[T]) Del(lat float64, lng float64) {
	r.index.Delete(geohash.EncodeInt(lat, lng))
}

func (r *GeohashBTree[T]) Get(lat float64, lng float64) (T, bool) {
	return r.index.Get(geohash.EncodeInt(lat, lng))
}

func (r *GeohashBTree[T]) Iter(iter IterFunc[T]) {
	r.index.Scan(func(key uint64, value T) bool {
		return iter(DecodeToPoint(key, value))
	})
}

type RangeSearchInput struct {
	minLat float64
	minLng float64
	maxLat float64
	maxLng float64
}

func (r *GeohashBTree[T]) RangeSearch(input *RangeSearchInput, iter IterFunc[T]) {
	r.index.Ascend(geohash.EncodeInt(input.minLat, input.minLng), func(key uint64, value T) bool {
		point := DecodeToPoint(key, value)

		if point.lat > input.maxLat || point.lng > input.maxLng {
			return false
		}

		return iter(point)
	})
}
