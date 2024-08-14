package qry

import (
	"github.com/reiver/go-erorr"
)

const (
	errNilEvaluator = erorr.Error("qry: nil evaluator")
	errNilMarshaler = erorr.Error("qry: nil marhshaler")
)
