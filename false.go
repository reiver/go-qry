package qry

type False[T any] struct {}

func (receiver False[T]) Evaluate(value T) (bool, error) {
	return false, nil
}
