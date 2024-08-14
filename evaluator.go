package qry

type Evaluator[T any] interface {
	Evaluate(T) (bool, error)
}
