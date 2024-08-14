package qry_test

import (
	"testing"

	"github.com/reiver/go-qry"
)

func TestTrue_string(t *testing.T) {

	tests := []struct{
		String string
		Expected bool
	}{
		{
			String: "",
			Expected: true,
		},



		{
			String: "apple",
			Expected: true,
		},
		{
			String: "banana",
			Expected: true,
		},
		{
			String: "cherry",
			Expected: true,
		},
	}

	for testNumber, test := range tests {

		var evaluator qry.Evaluator[string] = qry.True[string]{}

		actual, err := evaluator.Evaluate(test.String)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			continue
		}

		{
			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
				t.Logf("EXPECTED: %t", expected)
				t.Logf("ACTUAL:   %t", actual)
				continue
			}
		}
	}
}
