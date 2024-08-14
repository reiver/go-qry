package qry_test

import (
	"testing"

	"github.com/reiver/go-qry"
)

func TestNot_MarshalQRY_string(t *testing.T) {

	tests := []struct{
		Unit qry.Unit[string]
		Expected string
	}{
		{
			Unit: qry.False[string]{},
			Expected: "{not false}",
		},
		{
			Unit: qry.True[string]{},
			Expected: "{not true}",
		},
	}

	for testNumber, test := range tests {

		var marshaler qry.Marshaler = qry.Not[string]{Unit:test.Unit}

		actualBytes, err := qry.Marshal(marshaler)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("UNIT:  %#v", test.Unit)
			continue
		}

		{
			expected := test.Expected
			actual := string(actualBytes)

			if expected != actual {
				t.Errorf("For test #%d, the actual qry-marshaled value is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("UNIT:  %#v", test.Unit)
				continue
			}
		}
	}
}
