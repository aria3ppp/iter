package iter_test

import (
	"strconv"
	"testing"

	"github.com/aria3ppp/iter"
)

func TestFold(t *testing.T) {
	it := iter.FromSlice([]int{1, 2, 3, 4, 5})
	fldVal := iter.Fold(it, "", func(f string, item int) string {
		f += strconv.Itoa(item)
		return f
	})

	if fldVal != "12345" {
		t.Fatalf("expected folded value=%q, got=%q\n", "12345", fldVal)
	}
}
