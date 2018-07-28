package hashuri_test

import (
	"github.com/reiver/go-hashuri"

	"fmt"
	"net/url"
)

func ExampleType_parsequery() {

	var hashURI string = "hash://sha256/0ba904eae8773b70c75333db4de2f3ac45a8ad4ddba1b242f0b3cfc199391dd8?apple=one&banana=۲&cherry=3#over-there"

	var res hashuri.Type

	if err := hashuri.Parse(&res, hashURI); nil != err {
		panic(err)
		return
	}


	query, err := url.ParseQuery(res.RawQuery)
	if nil != err {
		panic(err)
		return
	}

	fmt.Printf("Hash URI raw query: %q\n", res.RawQuery)

	for i, key := range []string{"apple", "banana", "cherry"} {
		value := query.Get(key)

		fmt.Printf("key-value pair #%d: %q -> %q\n", 1+i, key, value)

	}

	// Output:
	// Hash URI raw query: "apple=one&banana=۲&cherry=3"
	// key-value pair #1: "apple" -> "one"
	// key-value pair #2: "banana" -> "۲"
	// key-value pair #3: "cherry" -> "3"
}
