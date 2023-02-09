package iter

func (it *Iterable[T]) Filter(fn func(T) bool) *Iterable[T] {
	fit := &filterIter[T]{inner: it.inner, fn: fn}
	return &Iterable[T]{inner: fit}
}

type filterIter[T any] struct {
	inner Iterator[T]
	fn    func(T) bool
}

func (fit *filterIter[T]) Next(v *T) bool {
	var innerVal T
	if fit.inner.Next(&innerVal) {
		if !fit.fn(innerVal) {
			return fit.Next(v)
		}
		*v = innerVal
		return true
	}
	return false
}