package qry

type Unit[T any] interface {
	Evaluator[T]
	Marshaler
}
