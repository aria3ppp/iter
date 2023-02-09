package iter_test

import (
	"testing"

	"github.com/aria3ppp/iter"
)

func TestReduce(t *testing.T) {
	it := iter.FromSlice([]int{1, 2, 3, 4, 5})
	rdcVal := iter.Reduce(it, func(r int, item int) int {
		r += item
		return r
	})

	if rdcVal != 15 {
		t.Errorf("expected reduced value=%d, got=%d\n", 15, rdcVal)
	}
}
