/*
Package hashuri parses Hash URIs.

And works seemlessly with the Go built-in packages: "database/sql", "fmt", "encoding/json".


Hash URI Examples

Hash URIs often look like this:

	 hash://sha256/0ba904eae8773b70c75333db4de2f3ac45a8ad4ddba1b242f0b3cfc199391dd8
	 \__/   \____/ \______________________________________________________________/
	  |       |                                    |
	scheme algorithm                              hash

But, more generally, can look like this:

	 hash://sha256/0ba904eae8773b70c75333db4de2f3ac45a8ad4ddba1b242f0b3cfc199391dd8?apple=one&banana=۲&cherry=3#over-there
	 \__/   \____/ \______________________________________________________________/ \_________________________/ \________/
	  |       |                                    |                                            |                   |
	scheme algorithm                              hash                                        query              fragment


Go Example

Basic usage of this package looks like this:

	import "github.com/reiver/go-hashuri"
	
	var hashURI string = "hash://sha256/0ba904eae8773b70c75333db4de2f3ac45a8ad4ddba1b242f0b3cfc199391dd8"
	
	var parsedHashURI hashuri.Type
	
	if err := hashuri.Parse(&parsedHashURI, hashURI); nil != err {
		return err
	}
	
	fmt.Printf("Hash URI algorithm: %q \n", parsedHashURI.Algorithm) // Hash URI algorithm: "sha256"
	fmt.Printf("Hash URI hash:      %q \n", parsedHashURI.Hash)      // Hash URI hash:      "0ba904eae8773b70c75333db4de2f3ac45a8ad4ddba1b242f0b3cfc199391dd8"
	fmt.Printf("Hash URI raw query: %q \n", parsedHashURI.RawQuery)  // Hash URI raw query: ""
	fmt.Printf("Hash URI fragment:  %q \n", parsedHashURI.Fragment)  // Hash URI fragment:  ""


See Also

The Hash URI specification can be found at: https://github.com/hash-uri/hash-uri
*/
package hashuri
