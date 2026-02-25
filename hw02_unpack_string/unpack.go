package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var (
	ErrInvalidString = errors.New("invalid string")
)

func Unpack(str string) (string, error) {
	var builder strings.Builder
	var prev rune
	var hasPrev, escaped bool

	for _, r := range str {
		if escaped {
			if !unicode.IsDigit(r) && r != '\\' {
				return "", ErrInvalidString
			}
			prev, hasPrev, escaped = r, true, false
			continue
		}

		if r == '\\' {
			if hasPrev {
				builder.WriteRune(prev)
			}
			escaped = true
			hasPrev = false
			continue
		}

		if unicode.IsDigit(r) {
			if !hasPrev {
				return "", ErrInvalidString
			}
			count := int(r - '0')
			builder.WriteString(strings.Repeat(string(prev), count))
			hasPrev = false
			continue
		}

		if hasPrev {
			builder.WriteRune(prev)
		}
		prev, hasPrev = r, true
	}

	if escaped {
		return "", ErrInvalidString
	}
	if hasPrev {
		builder.WriteRune(prev)
	}

	return builder.String(), nil
}
