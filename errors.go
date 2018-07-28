package hashuri

import (
	"errors"
)

var (
	errInternalError  = errors.New("Internal Error")
	errNilDestination = errors.New("Nil Destination")
	errNotHashURI     = errors.New("Not Hash URI")
)
