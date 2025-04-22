package parser

import (
	"fmt"
	"strconv"
	"strings"
	"test_alef_1_golang/internal/evaluator"
	"test_alef_1_golang/internal/utils"
)

func ParseExpression(input string) (evaluator.Evaluator, error) {
	input = strings.TrimSpace(input)

	if utils.IsNumeric(input) {
		value, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return nil, err
		}
		return &evaluator.Number{Value: value}, nil
	}

	if strings.HasPrefix(input, "(") && strings.HasSuffix(input, ")") {
		content := input[1 : len(input)-1]
		tokens, err := tokenize(content)
		if err != nil {
			return nil, err
		}

		operation := tokens[0]
		args := tokens[1:]

		parsedArgs := make([]evaluator.Evaluator, len(args))
		for i, arg := range args {
			parsedArg, err := ParseExpression(arg)
			if err != nil {
				return nil, err
			}
			parsedArgs[i] = parsedArg
		}

		return &evaluator.Operation{Operator: operation, Args: parsedArgs}, nil
	}

	return nil, fmt.Errorf("invalid expression: %s", input)
}
