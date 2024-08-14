package qry_test

import (
	"testing"

	"github.com/reiver/go-qry"
)

func TestOr_MarshalQRY_string(t *testing.T) {

	tests := []struct{
		Units []qry.Unit[string]
		Expected string
	}{
		{
			Units: []qry.Unit[string]{qry.False[string]{}},
			Expected: "{or false}",
		},
		{
			Units: []qry.Unit[string]{qry.True[string]{}},
			Expected: "{or true}",
		},



		{
			Units: []qry.Unit[string]{qry.False[string]{}, qry.False[string]{}},
			Expected: "{or false false}",
		},
		{
			Units: []qry.Unit[string]{qry.False[string]{}, qry.True[string]{}},
			Expected: "{or false true}",
		},
		{
			Units: []qry.Unit[string]{qry.True[string]{}, qry.False[string]{}},
			Expected: "{or true false}",
		},
		{
			Units: []qry.Unit[string]{qry.True[string]{}, qry.True[string]{}},
			Expected: "{or true true}",
		},



		{
			Units: []qry.Unit[string]{qry.False[string]{}, qry.False[string]{}, qry.False[string]{}},
			Expected: "{or false false false}",
		},
		{
			Units: []qry.Unit[string]{qry.False[string]{}, qry.False[string]{}, qry.True[string]{}},
			Expected: "{or false false true}",
		},
		{
			Units: []qry.Unit[string]{qry.False[string]{}, qry.True[string]{}, qry.False[string]{}},
			Expected: "{or false true false}",
		},
		{
			Units: []qry.Unit[string]{qry.False[string]{}, qry.True[string]{}, qry.True[string]{}},
			Expected: "{or false true true}",
		},
		{
			Units: []qry.Unit[string]{qry.True[string]{}, qry.False[string]{}, qry.False[string]{}},
			Expected: "{or true false false}",
		},
		{
			Units: []qry.Unit[string]{qry.True[string]{}, qry.False[string]{}, qry.True[string]{}},
			Expected: "{or true false true}",
		},
		{
			Units: []qry.Unit[string]{qry.True[string]{}, qry.True[string]{}, qry.False[string]{}},
			Expected: "{or true true false}",
		},
		{
			Units: []qry.Unit[string]{qry.True[string]{}, qry.True[string]{}, qry.True[string]{}},
			Expected: "{or true true true}",
		},
	}

	for testNumber, test := range tests {

		var marshaler qry.Marshaler = qry.Or[string]{Units:test.Units}

		actualBytes, err := qry.Marshal(marshaler)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("UNITS:  %#v", test.Units)
			continue
		}

		{
			expected := test.Expected
			actual := string(actualBytes)

			if expected != actual {
				t.Errorf("For test #%d, the actual qry-marshaled value is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("UNITS:  %#v", test.Units)
				continue
			}
		}
	}
}
