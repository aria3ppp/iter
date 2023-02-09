package iter

func Map[From any, To any](from *Iterable[From], fn func(From) To) *Iterable[To] {
	miter := &mapIter[From, To]{inner: from.inner, fn: fn}
	return &Iterable[To]{inner: miter}
}

/*
func (it *Iterable[From]) Map[Into any](fn func(From) Into) *Iterable[Into] {
	miter := &mapIter[From, Into]{inner: it.inner, fn: fn}
	return &Iterable[Into]{inner: miter}
}
*/

type mapIter[From any, To any] struct {
	inner Iterator[From]
	fn    func(From) To
}

func (mit *mapIter[From, To]) Next(v *To) bool {
	var innerVal From
	if mit.inner.Next(&innerVal) {
		*v = mit.fn(innerVal)
		return true
	}
	return false
}
