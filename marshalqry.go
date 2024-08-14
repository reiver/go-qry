package qry

import (
	"github.com/reiver/go-erorr"
)

func marshalQRY[T any](name string, units ...Unit[T]) ([]byte, error) {

	var buffer [256]byte
	var p []byte = buffer[0:0]

	p = append(p, '{')
	p = append(p, name...)

	for index, unit := range units {
		var marshaler Marshaler = unit

		if nil == marshaler {
			return nil, errNilMarshaler
		}

		p = append(p, ' ')

		marshaled, err := marshaler.MarshalQRY()
		if nil != err {
			var unitNumber int = index + 1
			return nil, erorr.Errorf("qry: problem marshaling sub-unit %d of an %q: %w", unitNumber, name, err)
		}

		p = append(p, marshaled...)
	}

	p = append(p, '}')

	return p, nil
}

