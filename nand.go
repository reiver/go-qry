package qry

type Nand[T any] struct {
	Evaluators []Evaluator[T]
}

var _ Evaluator[string] = Nand[string]{}

func (receiver Nand[T]) Evaluate(value T) (bool, error) {
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
			return true, nil
		}
	}
	return false, nil
}
