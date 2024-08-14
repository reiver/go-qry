package qry

type Not[T any] struct {
	Unit Unit[T]
}

var _ Unit[string] = Not[string]{}

func (receiver Not[T]) Evaluate(value T) (bool, error) {

	var empty bool

	var evaluator Evaluator[T] = receiver.Unit
	if nil == evaluator {
		return empty, errNilEvaluator
	}

	result, err := evaluator.Evaluate(value)
	if nil != err {
		return empty, err
	}

	return !result, nil
}

func (receiver Not[T]) MarshalQRY() ([]byte, error) {
	const name string = "not"
	return marshalQRY[T](name, receiver.Unit)
}
