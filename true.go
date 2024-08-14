package qry

type True[T any] struct {}

func (receiver True[T]) Evaluate(value T) (bool, error) {
	return true, nil
}
