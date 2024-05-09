package ddosml

type Request struct {
	IsDDOS bool
}

// data method returns slice of float64 representation of
// the Request properties (attributes)
//
//	prevData is a previous data rows, used to normalize values to floats.
func (c *Request) data(prevData []*Request) []float64 {
	//todo: normalize data and return []float64 based on that
	return nil
}

// isDDOSInt method returns isDDOSInt of given classifier model
// (0 if given req is DDOS, 1 if given request isn't DDOS)
func (c *Request) isDDOSInt() int {
	if c.IsDDOS {
		return 0
	}

	return 1
}
