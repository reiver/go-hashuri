package hashuri

import (
	"encoding/json"
)

// UnmarshalJSON is to make it so hashuri.Type fits the encoding/json.Unmarshaler interface.
func (receiver *Type) UnmarshalJSON(src []byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	var s string
	if err := json.Unmarshal(src, &s); nil != err {
		return err
	}

	if err := receiver.Scan(s); nil != err {
		return err
	}

	return nil
}
