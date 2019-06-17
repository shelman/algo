package main

import "fmt"
import "strconv"
import "strings"

func mustConvertTokenToInt(token string) int {
	conv, err := strconv.Atoi(strings.Trim(token, " "))
	if err != nil {
		panic(fmt.Sprintf("Can't convert %s to an int\n", token))
	}
	return conv
}

func evaluate(input string) int {
	accum := 0
	op := '+'
	idx := 0

	applyOp := func(val int) {
		switch op {
		case '-':
			accum = accum - val
		case '+':
			accum = accum + val
		}
	}

	getNextInt := func() int {
		iStr := ""
		for idx < len(input) && input[idx] != '-' && input[idx] != '+' && input[idx] != '(' && input[idx] != ')' {
			iStr = iStr + string(input[idx])
			idx++
		}
		return mustConvertTokenToInt(iStr)
	}

	for idx < len(input) {
		switch input[idx] {
		case ' ':
			idx++
		case '-', '+':
			op = rune(input[idx])
			idx++
		case '(':
			rParen := strings.Index(input[idx:], ")")
			applyOp(evaluate(input[idx+1:idx+rParen]))
			idx = idx+rParen+1
		default:
			applyOp(getNextInt())
		}
	}
	return accum
}

func main() {
	input := "6 + (2-1) +5"
	fmt.Println(evaluate(input))
}
