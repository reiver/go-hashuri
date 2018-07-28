package hashuri

import (
	"fmt"
)

// Scan is to make it so hashuri.Type fits the database/sql.Scanner interface.
func (receiver *Type) Scan(src interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	var value string
	{
		switch casted := src.(type) {
		case []byte:
			value = string(casted)

		case string:
			value = casted

		default:
			return fmt.Errorf("hashuri: cannot scan something that is not a string or []byte; namely: (%T) %v", src, src)
		}
	}

	if err := Parse(receiver, value); nil != err {
		return err
	}

	return nil
}
