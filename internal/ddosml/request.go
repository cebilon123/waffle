package ddosml

import "io"

// Request struct represents http request
// in the more "readable" format in order
// to work with that with ML.
type Request struct {
	IsDDOS           bool
	Method           string
	URL              string
	Protocol         string
	ProtocolMajor    int
	ProtocolMinor    int
	Headers          map[string]any
	Body             io.ReadCloser
	ContentLength    int64
	TransferEncoding []string
	Host             string
	Form             map[string][]string
	RemoteAddress    string
	RequestURI       string
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
