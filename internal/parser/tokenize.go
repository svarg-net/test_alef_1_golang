package parser

import "strings"

func tokenize(input string) ([]string, error) {
	var tokens []string
	var currentToken strings.Builder
	depth := 0

	for _, char := range input {
		switch char {
		case '(':
			depth++
		case ')':
			depth--
		case ' ':
			if depth == 0 {
				if currentToken.Len() > 0 {
					tokens = append(tokens, currentToken.String())
					currentToken.Reset()
				}
				continue
			}
		}
		currentToken.WriteRune(char)
	}

	if currentToken.Len() > 0 {
		tokens = append(tokens, currentToken.String())
	}

	return tokens, nil
}
