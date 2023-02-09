package iter

func Fold[T any, F any](it *Iterable[T], init F, fn func(F, T) F) F {
	accum := init
	var next T
	for it.inner.Next(&next) {
		accum = fn(accum, next)
	}
	return accum
}
