package main

import "fmt"
import pretty_print "github.com/k0kubun/pp"

func main() {
	var input_raw string
	fmt.Scan(&input_raw)
	pretty_print.Println()
}

func lex(input string) *Node {
	operator, index := find_operator(input)
	return &Node{
		Value: operator,
		Left:  &Node{Value: input[:index]},
		Right: lex(input[index+1:]),
	}
}

func find_operator(input string) (operator string, index int) {
	for i, char := range input {
		if string(char) == "+" || string(char) == "-" {
			return string(char), i
		}
	}
	return "", 0
}

func eval(tree Node) int {
	return 0
}

type Node struct {
	Value string
	Left  *Node
	Right *Node
}
