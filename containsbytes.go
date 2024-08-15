package qry

import (
	"bytes"
)

type ContainsBytes struct {
	Values [][]byte
}

var _ Unit[[]byte] = ContainsBytes{}

func (receiver ContainsBytes) Evaluate(text []byte) (bool, error) {
	for _, value := range receiver.Values {
		if !bytes.Contains(text, value) {
			return false, nil
		}
	}

        return true, nil
}

func (receiver ContainsBytes) MarshalQRY() ([]byte, error) {
	var buffer [256]byte
	var p []byte = buffer[0:0]

	p = append(p, "{contains"...)

	for _, value := range receiver.Values  {
		p = append(p, ' ')
		p = append(p, value...)
	}

	p = append(p, '}')

	return p, nil
}
