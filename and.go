package qry

type And[T any] struct {
	Units []Unit[T]
}

var _ Unit[string] = And[string]{}

func (receiver And[T]) Evaluate(value T) (bool, error) {
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

		if false == result {
			return false, nil
		}
	}
	return true, nil
}

func (receiver And[T]) MarshalQRY() ([]byte, error) {
	const name string = "and"
	return marshalQRY(name, receiver.Units...)
}
