package hashuri

// A hashuri.Type represents a parsed Hash URI.
//
// The general for represented is:
//
//
//	 hash://sha256/9f86d081884c7d659a2feaa0?type=text/plain#top
//	 \__/   \____/ \______________________/ \_____________/ \_/
//	  |       |               |                    |         |
//	Scheme Algorithm         Hash               RawQuery  Fragment
//
type Type struct {
	Algorithm string // hash function algorithm
	Hash      string // hash, encoded as hexadecimal
	RawQuery  string // encoded query values, without '?'
	Fragment  string // fragment, without '#'
}
