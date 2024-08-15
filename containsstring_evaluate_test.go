package qry_test

import (
	"testing"

	"github.com/reiver/go-qry"
)

func TestContainsString_Evaluate(t *testing.T) {

	tests := []struct{
		Unit qry.ContainsString
		Text string
		Expected bool
	}{
		{
			Unit: qry.ContainsString{},
			Text: "",
			Expected: true,
		},
		{
			Unit: qry.ContainsString{},
			Text: "apple",
			Expected: true,
		},
		{
			Unit: qry.ContainsString{},
			Text: "banana",
			Expected: true,
		},
		{
			Unit: qry.ContainsString{},
			Text: "cherry",
			Expected: true,
		},
		{
			Unit: qry.ContainsString{},
			Text: "Hello world!",
			Expected: true,
		},
		{
			Unit: qry.ContainsString{},
			Text: "😈",
			Expected: true,
		},
		{
			Unit: qry.ContainsString{},
			Text: "ONCE TWICE THRICE FOURCE",
			Expected: true,
		},



		{
			Unit: qry.ContainsString{"a"},
			Text: "",
			Expected: false,
		},
		{
			Unit: qry.ContainsString{"a"},
			Text: "apple",
			Expected: true,
		},
		{
			Unit: qry.ContainsString{"a"},
			Text: "banana",
			Expected: true,
		},
		{
			Unit: qry.ContainsString{"a"},
			Text: "cherry",
			Expected: false,
		},
		{
			Unit: qry.ContainsString{"a"},
			Text: "Hello world!",
			Expected: false,
		},
		{
			Unit: qry.ContainsString{"a"},
			Text: "😈",
			Expected: false,
		},
		{
			Unit: qry.ContainsString{"a"},
			Text: "ONCE TWICE THRICE FOURCE",
			Expected: false,
		},



		{
			Unit: qry.ContainsString{"banana"},
			Text: "",
			Expected: false,
		},
		{
			Unit: qry.ContainsString{"banana"},
			Text: "apple",
			Expected: false,
		},
		{
			Unit: qry.ContainsString{"banana"},
			Text: "banana",
			Expected: true,
		},
		{
			Unit: qry.ContainsString{"banana"},
			Text: "cherry",
			Expected: false,
		},
		{
			Unit: qry.ContainsString{"banana"},
			Text: "Hello world!",
			Expected: false,
		},
		{
			Unit: qry.ContainsString{"banana"},
			Text: "😈",
			Expected: false,
		},
		{
			Unit: qry.ContainsString{"banana"},
			Text: "ONCE TWICE THRICE FOURCE",
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
