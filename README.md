# go-hashuri

Package **hashuri** provides tools for working with **Hash URIs**, for the Go programming language.


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-hashuri

[![GoDoc](https://godoc.org/github.com/reiver/go-hashuri?status.svg)](https://godoc.org/github.com/reiver/go-hashuri)


## Example

```go
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
```


## See Also

The **Hash URI** specification can be found at: https://github.com/hash-uri/hash-uri
