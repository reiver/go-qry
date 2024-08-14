package qry_test

import (
	"testing"

	"github.com/reiver/go-qry"
)

func TestNand_Evaluate_string(t *testing.T) {

	tests := []struct{
		Units []qry.Unit[string]
		Expected bool
	}{
		{
			Units: []qry.Unit[string]{qry.False[string]{}},
			Expected: true,
		},
		{
			Units: []qry.Unit[string]{qry.True[string]{}},
			Expected: false,
		},



		{
			Units: []qry.Unit[string]{qry.False[string]{}, qry.False[string]{}},
			Expected: true,
		},
		{
			Units: []qry.Unit[string]{qry.False[string]{}, qry.True[string]{}},
			Expected: true,
		},
		{
			Units: []qry.Unit[string]{qry.True[string]{}, qry.False[string]{}},
			Expected: true,
		},
		{
			Units: []qry.Unit[string]{qry.True[string]{}, qry.True[string]{}},
			Expected: false,
		},



		{
			Units: []qry.Unit[string]{qry.False[string]{}, qry.False[string]{}, qry.False[string]{}},
			Expected: true,
		},
		{
			Units: []qry.Unit[string]{qry.False[string]{}, qry.False[string]{}, qry.True[string]{}},
			Expected: true,
		},
		{
			Units: []qry.Unit[string]{qry.False[string]{}, qry.True[string]{}, qry.False[string]{}},
			Expected: true,
		},
		{
			Units: []qry.Unit[string]{qry.False[string]{}, qry.True[string]{}, qry.True[string]{}},
			Expected: true,
		},
		{
			Units: []qry.Unit[string]{qry.True[string]{}, qry.False[string]{}, qry.False[string]{}},
			Expected: true,
		},
		{
			Units: []qry.Unit[string]{qry.True[string]{}, qry.False[string]{}, qry.True[string]{}},
			Expected: true,
		},
		{
			Units: []qry.Unit[string]{qry.True[string]{}, qry.True[string]{}, qry.False[string]{}},
			Expected: true,
		},
		{
			Units: []qry.Unit[string]{qry.True[string]{}, qry.True[string]{}, qry.True[string]{}},
			Expected: false,
		},
	}

	texts := []string{
		"",

		"apple",
		"banana",
		"cherry",

		"ONCE",
		"TWICE",
		"THRICE",
		"FOURCE",

		"😈",
	}

	for testNumber, test := range tests {

		for textNumber, text := range texts {

			var evaluator qry.Evaluator[string] = qry.Nand[string]{Units:test.Units}

			actual, err := evaluator.Evaluate(text)

			if nil != err {
				t.Errorf("For test #%d and text #%d, did not expect an error but actually got one.", testNumber, textNumber)
				t.Logf("ERROR: (%T) %s", err, err)
				t.Logf("TEXT: %q", text)
				t.Logf("UNIT:  %#v", test.Units)
				continue
			}

			{
				expected := test.Expected

				if expected != actual {
					t.Errorf("For test #%d and text #%d, the actual value is not what was expected.", testNumber, textNumber)
					t.Logf("EXPECTED: %t", expected)
					t.Logf("ACTUAL:   %t", actual)
					t.Logf("TEXT: %q", text)
					t.Logf("UNIT:  %#v", test.Units)
					continue
				}
			}
		}
	}
}
