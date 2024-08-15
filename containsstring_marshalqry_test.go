package qry_test

import (
	"testing"

	"github.com/reiver/go-qry"
)

func TestContainsString_MarshalQRY(t *testing.T) {

	tests := []struct{
		Unit qry.ContainsString
		Expected string
	}{
		{
			Unit: qry.ContainsString{Values:[]string{"apple"}},
			Expected:                      `{contains apple}`,
		},
		{
			Unit: qry.ContainsString{Values:[]string{"banana"}},
			Expected:                      `{contains banana}`,
		},
		{
			Unit: qry.ContainsString{Values:[]string{"cherry"}},
			Expected:                      `{contains cherry}`,
		},



		{
			Unit: qry.ContainsString{Values:[]string{"😈"}},
			Expected:                      `{contains 😈}`,
		},



		{
			Unit: qry.ContainsString{Values:[]string{"ONCE","TWICE","THRICE","FOURCE"}},
			Expected:                      `{contains ONCE TWICE THRICE FOURCE}`,
		},
	}

	for testNumber, test := range tests {

		actualBytes, err := qry.Marshal(test.Unit)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %q", err, err)
			t.Logf("UNIT: %#v", test.Unit)
			continue
		}

		{
			expected := test.Expected
			actual := string(actualBytes)

			if expected != actual {
				t.Errorf("For test #%d, the actual qry-marshaled value is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("UNIT: %#v", test.Unit)
				continue
			}
		}
	}
}
