package qry

import (
	"github.com/reiver/go-erorr"
)

const (
	errCannotMarshalNil = erorr.Error("qry: cannot qry-marshal nil")
	errNilEvaluator     = erorr.Error("qry: nil evaluator")
	errNilMarshaler     = erorr.Error("qry: nil marhshaler")
)
