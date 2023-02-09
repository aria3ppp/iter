package iter

func (it *Iterable[T]) ToSlice(sliceCap ...int) []T {
	cap := 0
	if len(sliceCap) != 0 {
		cap = sliceCap[0]
	}
	slc := make([]T, 0, cap)
	it.ForEach(func(v T) {
		slc = append(slc, v)
	})
	return slc
}