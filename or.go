package qry

type Or[T any] struct {
	Units []Unit[T]
}

var _ Unit[string] = Or[string]{}

func (receiver Or[T]) Evaluate(value T) (bool, error) {
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
			return true, nil
		}
	}
	return false, nil
}

func (receiver Or[T]) MarshalQRY() ([]byte, error) {
	const name string = "or"
	return marshalQRY(name, receiver.Units...)
}
