package qry

type False[T any] struct {}

var _ Unit[string] = False[string]{}

func (receiver False[T]) Evaluate(value T) (bool, error) {
	return false, nil
}

func (receiver False[T]) MarshalQRY() ([]byte, error) {
	var data [5]byte = [5]byte{'f','a','l','s','e'}
	return data[:], nil
}
