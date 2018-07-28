package hashuri

import (
	"encoding/json"
)

// MarshalJSON is to make it so hashuri.Type fits the encoding/json.Marshaler interface.
func (receiver Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(receiver.String())
}
