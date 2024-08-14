package qry

type Nor[T any] struct {
	Units []Unit[T]
}

var _ Unit[string] = Nor[string]{}

func (receiver Nor[T]) Evaluate(value T) (bool, error) {
	var empty bool

	for _, unit := range receiver.Units {
		var evaluator Evaluator[T] = unit

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

func (receiver Nor[T]) MarshalQRY() ([]byte, error) {
	const name string = "nor"
	return marshalQRY(name, receiver.Units...)
}
