package hashuri

import (
	"errors"
)

var (
	errInternalError  = errors.New("Internal Error")
	errNilDestination = errors.New("Nil Destination")
	errNilReceiver    = errors.New("Nil Receiver")
	errNotHashURI     = errors.New("Not Hash URI")
)
