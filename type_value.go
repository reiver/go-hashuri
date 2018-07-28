package hashuri

import (
	"database/sql/driver"
)

// Value is to make it so hashuri.Type fits the database/sql/driver.Valuer interface.
func (receiver Type) Value() (driver.Value, error) {
	s := receiver.String()
	return s, nil
}
