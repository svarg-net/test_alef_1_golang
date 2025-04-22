package evaluator

type Evaluator interface {
	Evaluate() (float64, error)
}
