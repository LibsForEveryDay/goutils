package stringutils

import (
	"strings"
)

func StartsWith(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

func EndsWith(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}
