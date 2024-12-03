package handlers

import (
	"strings"
	"unicode"
)

func IsBlankOrInvisible(s string) bool {
	trimmed := strings.TrimSpace(s)
	if len(trimmed) == 0 {
		return true
	}

	for _, r := range trimmed {
		if isInvisibleRune(r) {
			return true
		}
	}
	return false
}

func isInvisibleRune(r rune) bool {
	if unicode.IsSpace(r) || unicode.Is(unicode.Cf, r) || unicode.Is(unicode.Zl, r) || unicode.Is(unicode.Zp, r) {
		return true
	}

	invisibleRunes := []rune{
		'\u3164',
		'\u200B',
		'\u2060',
		'\uFEFF',
	}
	for _, invisibleRune := range invisibleRunes {
		if r == invisibleRune {
			return true
		}
	}
	return false
}
