package rule

type PredicateLogicalExpression func(payload []byte) (bool, error)

type Predicate struct {
	Name              string
	LogicalExpression PredicateLogicalExpression
}

type Builder interface {
	Build(name, variable, expression string) (*Predicate, error)
}

type PredicateBuilder struct {
}

var _ Builder = (*PredicateBuilder)(nil)

func (p *PredicateBuilder) Build(name, variable, expression string) (*Predicate, error) {
	//TODO implement me
	panic("implement me")
}
