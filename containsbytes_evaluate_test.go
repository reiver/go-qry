package qry_test

import (
	"testing"

	"github.com/reiver/go-qry"
)

func TestContainsBytes_Evaluate(t *testing.T) {

	tests := []struct{
		Unit qry.ContainsBytes
		Text []byte
		Expected bool
	}{
		{
			Unit: qry.ContainsBytes{},
			Text: []byte(""),
			Expected: true,
		},
		{
			Unit: qry.ContainsBytes{},
			Text: []byte("apple"),
			Expected: true,
		},
		{
			Unit: qry.ContainsBytes{},
			Text: []byte("banana"),
			Expected: true,
		},
		{
			Unit: qry.ContainsBytes{},
			Text: []byte("cherry"),
			Expected: true,
		},
		{
			Unit: qry.ContainsBytes{},
			Text: []byte("Hello world!"),
			Expected: true,
		},
		{
			Unit: qry.ContainsBytes{},
			Text: []byte("😈"),
			Expected: true,
		},
		{
			Unit: qry.ContainsBytes{},
			Text: []byte("ONCE TWICE THRICE FOURCE"),
			Expected: true,
		},



		{
			Unit: qry.ContainsBytes{[]byte("a")},
			Text: []byte(""),
			Expected: false,
		},
		{
			Unit: qry.ContainsBytes{[]byte("a")},
			Text: []byte("apple"),
			Expected: true,
		},
		{
			Unit: qry.ContainsBytes{[]byte("a")},
			Text: []byte("banana"),
			Expected: true,
		},
		{
			Unit: qry.ContainsBytes{[]byte("a")},
			Text: []byte("cherry"),
			Expected: false,
		},
		{
			Unit: qry.ContainsBytes{[]byte("a")},
			Text: []byte("Hello world!"),
			Expected: false,
		},
		{
			Unit: qry.ContainsBytes{[]byte("a")},
			Text: []byte("😈"),
			Expected: false,
		},
		{
			Unit: qry.ContainsBytes{[]byte("a")},
			Text: []byte("ONCE TWICE THRICE FOURCE"),
			Expected: false,
		},



		{
			Unit: qry.ContainsBytes{[]byte("banana")},
			Text: []byte(""),
			Expected: false,
		},
		{
			Unit: qry.ContainsBytes{[]byte("banana")},
			Text: []byte("apple"),
			Expected: false,
		},
		{
			Unit: qry.ContainsBytes{[]byte("banana")},
			Text: []byte("banana"),
			Expected: true,
		},
		{
			Unit: qry.ContainsBytes{[]byte("banana")},
			Text: []byte("cherry"),
			Expected: false,
		},
		{
			Unit: qry.ContainsBytes{[]byte("banana")},
			Text: []byte("Hello world!"),
			Expected: false,
		},
		{
			Unit: qry.ContainsBytes{[]byte("banana")},
			Text: []byte("😈"),
			Expected: false,
		},
		{
			Unit: qry.ContainsBytes{[]byte("banana")},
			Text: []byte("ONCE TWICE THRICE FOURCE"),
			Expected: false,
		},
	}

	for testNumber, test := range tests {

		actualResult, err := test.Unit.Evaluate(test.Text)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually hoy one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("TEXT: %q", test.Text)
			t.Logf("UNIT: %#v", test.Unit)
			continue
		}

		{
			expected := test.Expected
			actual := actualResult

			if expected != actual {
				t.Errorf("For test #%d, the actual evaluated-result it not what was expected.", testNumber)
				t.Logf("EXPECTED: %t", expected)
				t.Logf("ACTUAL:   %t", actual)
				t.Logf("TEXT: %q", test.Text)
				t.Logf("UNIT: %#v", test.Unit)
				continue
			}
		}
		_ = actualResult
	}
}
