package qry

type True[T any] struct {}

var _ Evaluator[string] = True[string]{}

func (receiver True[T]) Evaluate(value T) (bool, error) {
	return true, nil
}
