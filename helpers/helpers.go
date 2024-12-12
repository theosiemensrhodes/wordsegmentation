package helpers

import (
	"regexp"
	"strings"

	"github.com/kennygrant/sanitize"
)

// Clean a string from special characters, white spaces and
// transform upper case letters to lower case.
func CleanString(s string) string {
	s = sanitize.Accents(strings.ToLower(s))
	return regexp.MustCompile("[^a-zA-Z0-9]+").ReplaceAllString(s, "")
}
