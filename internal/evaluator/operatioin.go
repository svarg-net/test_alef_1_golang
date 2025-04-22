package evaluator

import "fmt"

type Operation struct {
	Operator string
	Args     []Evaluator
}

func (op *Operation) Evaluate() (float64, error) {
	var result float64
	var err error

	switch op.Operator {
	case "+":
		result = 0
		for _, arg := range op.Args {
			val, err := arg.Evaluate()
			if err != nil {
				return 0, err
			}
			result += val
		}
	case "-":
		if len(op.Args) == 1 {
			val, err := op.Args[0].Evaluate()
			if err != nil {
				return 0, err
			}
			result = -val
		} else {
			result, err = op.Args[0].Evaluate()
			if err != nil {
				return 0, err
			}
			for _, arg := range op.Args[1:] {
				val, err := arg.Evaluate()
				if err != nil {
					return 0, err
				}
				result -= val
			}
		}
	case "*":
		result = 1
		for _, arg := range op.Args {
			val, err := arg.Evaluate()
			if err != nil {
				return 0, err
			}
			result *= val
		}
	case "/":
		result, err = op.Args[0].Evaluate()
		if err != nil {
			return 0, err
		}
		for _, arg := range op.Args[1:] {
			val, err := arg.Evaluate()
			if err != nil {
				return 0, err
			}
			if val == 0 {
				return 0, fmt.Errorf("division by zero")
			}
			result /= val
		}
	default:
		return 0, fmt.Errorf("unknown operation: %s", op.Operator)
	}

	return result, nil
}
