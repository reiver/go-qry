package qry

import (
	"github.com/reiver/go-erorr"
)

func Marshal(value any) ([]byte, error) {
	if nil == value {
		return nil, errCannotMarshalNil
	}

	switch casted := value.(type) {
	case Marshaler:
		return casted.MarshalQRY()
	default:
		return nil, erorr.Errorf("qry: cannot qry-marshal something of type %T", value)
	}
}
