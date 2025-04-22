package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"test_alef_1_golang/internal/parser"
)

func main() {
	var input string

	// Если аргумент передан через командную строку
	if len(os.Args) > 1 {
		input = os.Args[1]
	} else {
		fmt.Print("Введите выражение (например, (* (+ 1 3) 2)): ")
		reader := bufio.NewReader(os.Stdin)
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)
	}

	parsed, err := parser.ParseExpression(input)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	result, err := parsed.Evaluate()
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	fmt.Printf("Пример: %s\n", input)
	fmt.Printf("Результат: %.2f\n", result)
}
