package qry

type Nand[T any] struct {
	Units []Unit[T]
}

var _ Unit[string] = Nand[string]{}

func (receiver Nand[T]) Evaluate(value T) (bool, error) {
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
			return true, nil
		}
	}
	return false, nil
}

func (receiver Nand[T]) MarshalQRY() ([]byte, error) {
	const name string = "nand"
	return marshalQRY(name, receiver.Units...)
}
