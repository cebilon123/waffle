package ddosml

// DDOS represents machine learning
// DDOS protection.
type DDOS struct {
	isEnabled bool
}

// NewDDOS creates new ddos ML analyzer used to
// analyze requests in order to find out if given
// request is ddos attack or not.
func NewDDOS(isEnabled bool) *DDOS {
	return &DDOS{isEnabled: isEnabled}
}
