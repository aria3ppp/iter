package iter

func (it *Iterable[T]) Until(fn func(T) bool) *Iterable[T] {
	uit := &untilIter[T]{inner: it.inner, fn: fn}
	return &Iterable[T]{inner: uit}
}

type untilIter[T any] struct {
	inner Iterator[T]
	fn    func(T) bool
}

func (uit *untilIter[T]) Next(v *T) bool {
	var innerVal T
	if uit.inner.Next(&innerVal) {
		if !uit.fn(innerVal) {
			return false
		}
		*v = innerVal
		return true
	}
	return false
}