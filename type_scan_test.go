package hashuri

import (
	"testing"
)

func TestTypeScan(t *testing.T) {

	tests := []struct{
		Value    string
		Expected Type
	}{
		{
			Value: "hash://crunch/fedcba9876543210",
			Expected: Type{
				Algorithm: "crunch",
				Hash:      "fedcba9876543210",
				RawQuery:  "",
				Fragment:  "",
			},
		},
		{
			Value: "hash://crunch/fedcba9876543210?apple=one&banana=۲&cherry=3",
			Expected: Type{
				Algorithm: "crunch",
				Hash:      "fedcba9876543210",
				RawQuery:  "apple=one&banana=۲&cherry=3",
				Fragment:  "",
			},
		},
		{
			Value: "hash://crunch/fedcba9876543210#over-there",
			Expected: Type{
				Algorithm: "crunch",
				Hash:      "fedcba9876543210",
				RawQuery:  "",
				Fragment:  "over-there",
			},
		},
		{
			Value: "hash://crunch/fedcba9876543210?apple=one&banana=۲&cherry=3#over-there",
			Expected: Type{
				Algorithm: "crunch",
				Hash:      "fedcba9876543210",
				RawQuery:  "apple=one&banana=۲&cherry=3",
				Fragment:  "over-there",
			},
		},



		{
			Value: "hash://کوچک شدن/fedcba9876543210?apple=one&banana=۲&cherry=3#آنجا",
			Expected: Type{
				Algorithm: "کوچک شدن",
				Hash:      "fedcba9876543210",
				RawQuery:  "apple=one&banana=۲&cherry=3",
				Fragment:  "آنجا",
			},
		},
	}


	for testNumber, test := range tests {

		{
			var actualType Type

			actualType.Algorithm = "ERROR-Algorithm"
			actualType.Hash      = "ERROR-Hash"
			actualType.RawQuery  = "ERROR-RawQuery"
			actualType.Fragment  = "ERROR-Fragment"

			if err := actualType.Scan(test.Value); nil != err {
					t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
				t.Errorf("\tvalue:  %q", test.Value)
				t.Errorf("\tactual: %q", actualType.String())
				continue
			}

			if expected, actual := test.Expected.Algorithm, actualType.Algorithm; expected != actual {
				t.Errorf("For test #%d, algorithm expected %q, but actually got %q.", testNumber, expected, actual)
				continue
			}
			if expected, actual := test.Expected.Hash, actualType.Hash; expected != actual {
				t.Errorf("For test #%d, hash expected %q, but actually got %q.", testNumber, expected, actual)
				continue
			}
			if expected, actual := test.Expected.RawQuery, actualType.RawQuery; expected != actual {
				t.Errorf("For test #%d, raw query expected %q, but actually got %q.", testNumber, expected, actual)
				continue
			}
			if expected, actual := test.Expected.Fragment, actualType.Fragment; expected != actual {
				t.Errorf("For test #%d, fragment expected %q, but actually got %q.", testNumber, expected, actual)
				continue
			}
		}

		{
			var actualType Type

			actualType.Algorithm = "ERROR-Algorithm"
			actualType.Hash      = "ERROR-Hash"
			actualType.RawQuery  = "ERROR-RawQuery"
			actualType.Fragment  = "ERROR-Fragment"

			var bytes []byte = []byte(test.Value)

			if err := actualType.Scan(bytes); nil != err {
				t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
				t.Errorf("\tvalue:  %q", test.Value)
				t.Errorf("\tactual: %q", actualType.String())
				continue
			}

			if expected, actual := test.Expected.Algorithm, actualType.Algorithm; expected != actual {
				t.Errorf("For test #%d, algorithm expected %q, but actually got %q.", testNumber, expected, actual)
				continue
			}
			if expected, actual := test.Expected.Hash, actualType.Hash; expected != actual {
				t.Errorf("For test #%d, hash expected %q, but actually got %q.", testNumber, expected, actual)
				continue
			}
			if expected, actual := test.Expected.RawQuery, actualType.RawQuery; expected != actual {
				t.Errorf("For test #%d, raw query expected %q, but actually got %q.", testNumber, expected, actual)
				continue
			}
			if expected, actual := test.Expected.Fragment, actualType.Fragment; expected != actual {
				t.Errorf("For test #%d, fragment expected %q, but actually got %q.", testNumber, expected, actual)
				continue
			}
		}

		{
			var actualType Type

			actualType.Algorithm = "ERROR-Algorithm"
			actualType.Hash      = "ERROR-Hash"
			actualType.RawQuery  = "ERROR-RawQuery"
			actualType.Fragment  = "ERROR-Fragment"

			if err := actualType.Scan(5); nil == err {
				t.Errorf("For test #%d, expected an error, but did not actually get one: (%T) %q", testNumber, err, err)
				continue
			}
		}
	}
}
