package iter

func Reduce[T any](it *Iterable[T], fn func(T, T) T) T {
	var initVal T
	return Fold(it, initVal, fn)
}
