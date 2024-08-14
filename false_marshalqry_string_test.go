package qry_test

import (
	"testing"

	"github.com/reiver/go-qry"
)

func TestFalse_MarshalQRY(t *testing.T) {

	var marshaler qry.Marshaler = qry.False[string]{}

	actualBytes, err := qry.Marshal(marshaler)

	if nil != err {
		t.Error("Did not expect to get an error but actually got one.")
		t.Logf("ERROR: (%T) %s", err, err)
		return
	}

	{
		var expected string = "false"
		var actual   string = string(actualBytes)

		if expected != actual {
			t.Error("The actual marshaled value is not what was expected.")
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			return
		}
	}

}
