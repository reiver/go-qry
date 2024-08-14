package qry

type Nor[T any] struct {
	Evaluators []Evaluator[T]
}

var _ Evaluator[string] = Nor[string]{}

func (receiver Nor[T]) Evaluate(value T) (bool, error) {
	var empty bool

	for _, evaluator := range receiver.Evaluators {
		if nil == evaluator {
			return empty, errNilEvaluator
		}

		result, err := evaluator.Evaluate(value)
		if nil != err {
			return empty, err
		}

		if true == result {
			return false, nil
		}
	}
	return true, nil
}
