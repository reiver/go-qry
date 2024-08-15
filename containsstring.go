package qry

import (
	"strings"
)

type ContainsString struct {
	Values []string
}

var _ Unit[string] = ContainsString{}

func (receiver ContainsString) Evaluate(text string) (bool, error) {
	for _, value := range receiver.Values {
		if !strings.Contains(text, value) {
			return false, nil
		}
	}

	return true, nil
}

func (receiver ContainsString) MarshalQRY() ([]byte, error) {
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

