// Package xss can be used to validate input against set of xss rules.
package xss

import (
	"github.com/corazawaf/libinjection-go"
)

// IsXSS returns true whenever input string contains XSS.
func IsXSS(input string) bool {
	return libinjection.IsXSS(input)
}
