package spacial_queries

import (
	"fmt"
	"github.com/tidwall/btree"
	"testing"
)

func TestUnit_GeohashBtree_Ok(t *testing.T) {
	map_ := btree.Map[int64, string]{}
	map_.Set(10, "Hi")
	map_.Set(20, "Hello")

	for _, key := range map_.Keys() {
		v, _ := map_.Get(key)
		fmt.Printf("%d: %s\n", key, v)
	}
}
