package qry

type True[T any] struct {}

var _ Unit[string] = True[string]{}

func (receiver True[T]) Evaluate(value T) (bool, error) {
	return true, nil
}

func (receiver True[T]) MarshalQRY() ([]byte, error) {
	var data [6]byte = [6]byte{'{','t','r','u','e','}'}
	return data[:], nil
}
