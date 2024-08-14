package qry

type Identity[T any] struct {
	Evaluator Evaluator[T]
}

var _ Evaluator[string] = Identity[string]{}

func (receiver Identity[T]) Evaluate(value T) (bool, error) {

	var empty bool

	var evaluator Evaluator[T] = receiver.Evaluator
	if nil == evaluator {
		return empty, errNilEvaluator
	}

	result, err := evaluator.Evaluate(value)
	if nil != err {
		return empty, err
	}

	return result, nil
}
