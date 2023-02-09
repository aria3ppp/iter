package iter

func (it *Iterable[T]) ForEach(fn func(T)) {
	var v T
	for it.inner.Next(&v) {
		fn(v)
	}
}