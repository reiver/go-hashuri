package hashuri

import (
	"bytes"
)

// String reassembles the Hash URI into a valid Hash URI string. The general form of the result is:
//
//	 hash://sha256/9f86d081884c7d659a2feaa0?type=text/plain#top
//	 \__/   \____/ \______________________/ \_____________/ \_/
//	  |       |               |                    |         |
//	scheme algorithm         hash                query    fragment
func (receiver Type) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("hash:")

	if "" != receiver.Algorithm {
		buffer.WriteString("//")
		buffer.WriteString(receiver.Algorithm)
	}

	if "" != receiver.Algorithm || "" != receiver.Hash {
		buffer.WriteRune('/')
	}

	if "" != receiver.Hash {
		buffer.WriteString(receiver.Hash)
	}

	if "" != receiver.RawQuery {
		buffer.WriteRune('?')
		buffer.WriteString(receiver.RawQuery)
	}
	if "" != receiver.Fragment {
		buffer.WriteRune('#')
		buffer.WriteString(receiver.Fragment)
	}

	return buffer.String()
}
