package rule

const (
	operatorEq    = "=="
	operatorNotEq = "!="
	operatorGt    = ">"
	operatorLs    = "<"
	operatorArrow = "=>"

	methodLength = "LEN"
	methodFormat = "FORMAT"

	fieldPayload = "payload"
	fieldHeaders = "headers"

	tokenVariable = "variable"
)

var (
	// operators are compilable operators.
	operators = []string{
		operatorEq,
		operatorNotEq,
		operatorGt,
		operatorLs,
		operatorArrow,
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

type Tokenizer interface {
	BuildTokens(variable, expression string) ([]Token, error)
}

type Token struct {
	Name  string
	Value string
}

type tokenizer struct {
}

func (t *tokenizer) BuildTokens(variable, expression string) ([]Token, error) {
	for _, r := range []rune(expression) {

	}
}
