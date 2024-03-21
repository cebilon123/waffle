package rule

import "waffle/internal/request"

type node interface {
	Eval(r request.Wrapper) bool
}

type and struct {
	left, right node
}

func (a and) Eval(r request.Wrapper) bool {
	return a.left.Eval(r) && a.right.Eval(r)
}

func (a and) SetChild(left, right node) {
	a.left = left
	a.right = right
}

type or struct {
	left, right node
}

func (o or) Eval(r request.Wrapper) bool {
	return o.left.Eval(r) || o.right.Eval(r)
}

func (o or) SetChild(left, right node) {
	o.left = left
	o.right = right
}

type eq struct {
	left, right node
}

func (e eq) Eval(r request.Wrapper) bool {
	return e.left.Eval(r) == e.right.Eval(r)
}

type gt struct {
	valueFunc func(r request.Wrapper) (int, error)
	check     int
}

func (g gt) Eval(r request.Wrapper) bool {
	value, err := g.valueFunc(r)
	if err != nil {
		return false
	}

	return value > g.check
}
