package qry

import (
	"strconv"
	"bytes"
)

type ContainsBytes struct {
	Text []byte
}

var _ Unit[[]byte] = ContainsBytes{}

func (receiver ContainsBytes) Evaluate(value []byte) (bool, error) {
	return bytes.Contains(value, receiver.Text), nil
}

func (receiver ContainsBytes) MarshalQRY() ([]byte, error) {
	var buffer [256]byte
	var p []byte = buffer[0:0]

	p = append(p, "{contains "...)
	p = append(p, strconv.Quote(string(receiver.Text))...)
	p = append(p, '}')

	return p, nil
}
