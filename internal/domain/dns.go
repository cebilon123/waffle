// Package domain contains all the code and logic
// related to working with the internal domain system
package domain

// NameSystemProvider provides address of the destination for given
// registered domain.
type NameSystemProvider interface {
	GetAddress(host string) (string, error)
}
