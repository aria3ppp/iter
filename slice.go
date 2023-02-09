package iter

func FromValues[T any](vals ...T) *Iterable[T] {
	if len(vals) == 0 {
		panic("iter.FromValues: no value supplied")
	}
	return FromSlice(vals)
}

func FromSlice[T any](slc []T) *Iterable[T] {
	iterator := &sliceIter[T]{slc: slc}
	return NewIterable[T](iterator)
}

type sliceIter[T any] struct {
	slc      []T
	curIndex int
}

func (si *sliceIter[T]) Next(v *T) bool {
	if si.curIndex < len(si.slc) {
		*v = si.slc[si.curIndex]
		si.curIndex++
		return true
	}
	return false
}
