package evaluator

type Number struct {
	Value float64
}

func (n *Number) Evaluate() (float64, error) {
	return n.Value, nil
}
