package qry

import (
	"strconv"
	"strings"
)

type ContainsString struct {
	Text string
}

var _ Unit[string] = ContainsString{}

func (receiver ContainsString) Evaluate(value string) (bool, error) {
	return strings.Contains(value, receiver.Text), nil
}

func (receiver ContainsString) MarshalQRY() ([]byte, error) {
	var buffer [256]byte
	var p []byte = buffer[0:0]

	p = append(p, "{contains "...)
	p = append(p, strconv.Quote(receiver.Text)...)
	p = append(p, '}')

	return p, nil
}
