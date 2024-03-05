// Package domain contains all the code and logic
// related to working with the internal domain system
package domain

import "net/url"

// NameSystemProvider provides address of the destination for given
// registered domain.
type NameSystemProvider interface {
	GetAddress(host string) (*url.URL, error)
}
