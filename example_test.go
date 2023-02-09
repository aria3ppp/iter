package iter_test

import (
	"fmt"
	"strconv"

	"github.com/aria3ppp/iter"
)

func ExampleMap() {
	it := iter.FromValues("5", "2", "7", "12", "6")

	iter.Map(it, func(item string) int {
		item += item
		n, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}
		return n
	}).
		ForEach(func(n int) {
			fmt.Printf("%d\n", n)
		})

	// Output:
	// 55
	// 22
	// 77
	// 1212
	// 66
}

func ExampleFilter() {
	iter.FromValues(5, 2, 7, 12, 6).
		Filter(func(n int) bool {
			return n > 6
		}).
		ForEach(func(n int) {
			fmt.Println(n)
		})

	// Output:
	// 7
	// 12
}

func ExampleUntil() {
	iter.FromValues(1, 2, 3, 4, 5, 6, 7).
		Until(func(n int) bool {
			return n < 4
		}).
		ForEach(func(n int) {
			fmt.Println(n)
		})

	// Output:
	// 1
	// 2
	// 3
}
