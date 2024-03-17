package rule

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

type expressionTree node

type Tokenizer interface {
	BuildExpressionTree(variable, expression string) (expressionTree, error)
}

type customRulesTokenizer struct {
}

var _ Tokenizer = (*customRulesTokenizer)(nil)

func newCustomRulesTokenizer() *customRulesTokenizer {
	return &customRulesTokenizer{}
}

func (c *customRulesTokenizer) BuildExpressionTree(variable, expression string) (expressionTree, error) {
	//TODO implement me
	panic("implement me")
}
