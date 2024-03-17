package rule

import (
	"waffle/internal/request"
)

const (
	operatorEq    = "=="
	operatorNotEq = "!="
	operatorGt    = ">"
	operatorLs    = "<"

	methodLength = "LEN"
	methodFormat = "FORMAT"

	fieldPayload = "payload"
	fieldHeaders = "headers"
)

var (
	// operators are compilable operators.
	operators = []string{
		operatorEq,
		operatorNotEq,
		operatorGt,
		operatorLs,
	}

	// methods are compilable methods.
	methods = []string{
		methodLength,
		methodFormat,
	}

	// fields are compilable fields available in the request wrapper.
	fields = []string{
		fieldPayload,
		fieldHeaders,
	}
)

type Predicate struct {
	Name string
	Eval func(r *request.Wrapper) bool
}

type Builder interface {
	Build(name, variable, expression string) (*Predicate, error)
}

type predicateBuilder struct {
}

var _ Builder = (*predicateBuilder)(nil)

func (p *predicateBuilder) Build(name, variable, expression string) (*Predicate, error) {
	//tree := and{
	//	left: gt{
	//		valueFunc: func(r request.Wrapper) (int, error) {
	//			bytes, err := io.ReadAll(r.Request.Body)
	//			if err != nil {
	//				return 0, fmt.Errorf("cannot read bytes: %w", err)
	//			}
	//
	//			return len(bytes), nil
	//		},
	//		check: 0,
	//	},
	//	right: gt{
	//		valueFunc: func(r request.Wrapper) (int, error) {
	//			return len(r.Request.Header), nil
	//		},
	//		check: 0,
	//	},
	//}
	return nil, nil
}
