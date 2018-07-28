package hashuri_test

import (
	"github.com/reiver/go-hashuri"

	"fmt"
)

func ExampleType_String() {

	var res hashuri.Type

	res.Algorithm = "sha256"
	res.Hash      = "0ba904eae8773b70c75333db4de2f3ac45a8ad4ddba1b242f0b3cfc199391dd8"
	res.RawQuery  = "apple=one&banana=۲&cherry=3"
	res.Fragment  = "over-there"

	fmt.Println(res.String())

	// Output:
	// hash://sha256/0ba904eae8773b70c75333db4de2f3ac45a8ad4ddba1b242f0b3cfc199391dd8?apple=one&banana=۲&cherry=3#over-there
}
