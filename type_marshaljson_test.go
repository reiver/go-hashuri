package hashuri

import (
	"testing"
)

func TestTypeMarshalJSON(t *testing.T) {

	tests := []struct{
		Value    Type
		Expected string
	}{
		{
			Value: Type{},
			Expected: `"hash:"`,
		},



		{
			Value: Type{
				Algorithm: "md5",
			},
			Expected: `"hash://md5/"`,
		},
		{
			Value: Type{
				Algorithm: "sha1",
			},
			Expected: `"hash://sha1/"`,
		},
		{
			Value: Type{
				Algorithm: "sha256",
			},
			Expected: `"hash://sha256/"`,
		},



		{
			Value: Type{
				Hash: "0ba904eae8773b70c75333db4de2f3ac45a8ad4ddba1b242f0b3cfc199391dd8",
			},
			Expected: `"hash:/0ba904eae8773b70c75333db4de2f3ac45a8ad4ddba1b242f0b3cfc199391dd8"`,
		},



		{
			Value: Type{
				RawQuery: "apple=one&banana=۲&cherry=3",
			},
			Expected: `"hash:?apple=one\u0026banana=۲\u0026cherry=3"`,
		},



		{
			Value: Type{
				Fragment: "over-there",
			},
			Expected: `"hash:#over-there"`,
		},



		{
			Value: Type{
				Algorithm: "sha256",
				Hash: "0ba904eae8773b70c75333db4de2f3ac45a8ad4ddba1b242f0b3cfc199391dd8",
				RawQuery: "apple=one&banana=۲&cherry=3",
				Fragment: "over-there",
			},
			Expected: `"hash://sha256/0ba904eae8773b70c75333db4de2f3ac45a8ad4ddba1b242f0b3cfc199391dd8?apple=one\u0026banana=۲\u0026cherry=3#over-there"`,
		},



		{
			Value: Type{
				Algorithm: "sha256",
				Hash: "0ba904eae8773b70c75333db4de2f3ac45a8ad4ddba1b242f0b3cfc199391dd8",
				RawQuery: "apple=one&banana=۲&cherry=3",
			},
			Expected: `"hash://sha256/0ba904eae8773b70c75333db4de2f3ac45a8ad4ddba1b242f0b3cfc199391dd8?apple=one\u0026banana=۲\u0026cherry=3"`,
		},
		{
			Value: Type{
				Algorithm: "sha256",
				Hash: "0ba904eae8773b70c75333db4de2f3ac45a8ad4ddba1b242f0b3cfc199391dd8",
				Fragment: "over-there",
			},
			Expected: `"hash://sha256/0ba904eae8773b70c75333db4de2f3ac45a8ad4ddba1b242f0b3cfc199391dd8#over-there"`,
		},
		{
			Value: Type{
				Algorithm: "sha256",
				Hash: "0ba904eae8773b70c75333db4de2f3ac45a8ad4ddba1b242f0b3cfc199391dd8",
			},
			Expected: `"hash://sha256/0ba904eae8773b70c75333db4de2f3ac45a8ad4ddba1b242f0b3cfc199391dd8"`,
		},
	}


	for testNumber, test := range tests {

		actualBytes, err := test.Value.MarshalJSON()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
			continue
		}

		actual := string(actualBytes)

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d,...", testNumber)
			t.Errorf("\texpected: %q", expected)
			t.Errorf("\tactual:   %q", actual)
			continue
		}
	}
}
