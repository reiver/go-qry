package qry

type False[T any] struct {}

var _ Evaluator[string] = False[string]{}

func (receiver False[T]) Evaluate(value T) (bool, error) {
	return false, nil
}
