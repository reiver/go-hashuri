package hashuri

import (
	"strings"
)

// hashuri.Parse parses value into a hashuri.Type structure.
func Parse(dest *Type, value string) error {
	if nil == dest {
		return errNilDestination
	}

	s := value

	{
		const prefix string = "hash:"

		if len(s) < len(prefix) {
			return errNotHashURI
		}

		if !strings.HasPrefix(s, prefix) {
			return errNotHashURI
		}

		s = s[len(prefix):]
	}

	{
		const prefix string = "//"

		if !strings.HasPrefix(s, prefix) {
			return errNotHashURI
		}

		s = s[len(prefix):]
	}

	{
		index := strings.IndexRune(s, '/')
		if 0 > index {
			return errNotHashURI
		}

		dest.Algorithm = s[:index]

		s = s[1+index:]
	}

	{
		var index int = -1
		if 0 > index {
			index = strings.IndexRune(s, '?')
		}
		if 0 >  index {
			index = strings.IndexRune(s, '#')
		}

		switch {
		case 0 > index:
			dest.Hash = s
			return nil
		default:
			dest.Hash = s[:index]
		}

		s = s[index:]
	}

	{
		s0 := s[0]
		if '?' != s0 {
			goto skipRawQuery
		}

		index := strings.IndexRune(s, '#')

		switch {
		case 0 > index:
			dest.RawQuery = s[1:]
			return nil
		default:
			dest.RawQuery = s[1:index]
		}

		s = s[index:]
	}
	skipRawQuery:

	{
		s0 := s[0]
		if '#' != s0 {
			return errInternalError
		}

		dest.Fragment = s[1:]
	}

	return nil
}
