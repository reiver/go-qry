package qry

type And[T any] struct {
	Evaluators []Evaluator[T]
}

var _ Evaluator[string] = And[string]{}

func (receiver And[T]) Evaluate(value T) (bool, error) {
	var empty bool

	for _, evaluator := range receiver.Evaluators {
		if nil == evaluator {
			return empty, errNilEvaluator
		}

		result, err := evaluator.Evaluate(value)
		if nil != err {
			return empty, err
		}

		if false == result {
			return false, nil
		}
	}
	return true, nil
}
