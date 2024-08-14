package qry

type Identity[T any] struct {
	Unit Unit[T]
}

var _ Unit[string] = Identity[string]{}

func (receiver Identity[T]) Evaluate(value T) (bool, error) {

	var empty bool

	var evaluator Evaluator[T] = receiver.Unit
	if nil == evaluator {
		return empty, errNilEvaluator
	}

	result, err := evaluator.Evaluate(value)
	if nil != err {
		return empty, err
	}

	return result, nil
}

func (receiver Identity[T]) MarshalQRY() ([]byte, error) {
	const name string = "identity"
	return marshalQRY[T](name, receiver.Unit)
}
