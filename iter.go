package iter

// `Iterator`` is the building block to make an `Iterable`.
// Users just need to implement ‍`Iterator` interface
// and pass the object to `NewIterable` function
type Iterator[T any] interface {
	Next(*T) bool
}

// `Iterable` add some utility methods on top of `Iterator`
type Iterable[T any] struct {
	inner Iterator[T]
}

func NewIterable[T any](iterator Iterator[T]) *Iterable[T] {
	return &Iterable[T]{inner: iterator}
}
