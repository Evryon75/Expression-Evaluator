package main

import (
	"fmt"
	"strconv"
	"strings"
)
import pretty_print "github.com/k0kubun/pp" // pretty print libs my beloved :)

// this constant is treated like an operator for logic purposes, its value is arbitrary
const EOF string = "X"

func main() {
	println("Input your mathematical expression, input 'q' to quit")
	for true {
		var input_raw string
		fmt.Scan(&input_raw)
		if input_raw == "q" {
			break
		} else if strings.ContainsAny(input_raw, "(") || strings.ContainsAny(input_raw, ")") {
			fmt.Println("Grouping evaluation such as '(x+y)*z' is not implemented yet")
		} else {
			pretty_print.Println(eval(lex(EOF + input_raw)))
		}
	}
}

func lex(input string) *Node {
	operator, index := find_operator(input)
	if index > -1 {
		return &Node{
			Value: operator,
			Left:  &Node{Value: input[index+1:]},
			Right: lex(input[:index]),
		}
	} else {
		return &Node{Value: strconv.FormatFloat(eval_md(input), 'f', -1, 64)}
	}
}

func find_operator(input string) (operator string, index int) {
	// reverse for loop because we want to assemble the tree from right to left, otherwise it would have the opposite order
	for i := len(input) - 1; i >= 0; i-- {
		if string(input[i]) == "+" || string(input[i]) == "-" || string(input[i]) == EOF {
			return string(input[i]), i
		}
	}
	return "No plus or minus operators found", -1
}

func eval(tree *Node) float64 {
	var result float64 = 0
	switch tree.Value {
	case "+":
		if tree.Right.Value == EOF {
			result = eval(tree.Right.Left) + eval(tree.Left)
		} else {
			result = eval(tree.Right) + eval(tree.Left)
		}
	case "-":
		if tree.Right.Value == EOF {
			result = eval(tree.Right.Left) - eval(tree.Left)
		} else {
			result = eval(tree.Right) - eval(tree.Left)
		}
	default:
		if tree.Value == EOF {
			result = eval_md(strings.Replace(tree.Left.Value, EOF, "", 1))
		} else {
			result = eval_md(strings.Replace(tree.Value, EOF, "", 1))
		}
	}
	return result
}

func eval_md(input string) float64 {
	input_clean := []float64{}
	input_split := strings.Split(input, "*")
	for _, piece_m := range input_split {
		for i, piece_d := range strings.Split(piece_m, "/") {
			temp, _ := strconv.ParseFloat(piece_d, 64)
			if i == 0 {
				input_clean = append(input_clean, float64(temp))
			} else {
				input_clean = append(input_clean, 1/float64(temp))
			}
		}
	}
	var multiplicand float64 = 1.0
	for _, multiplicator := range input_clean {
		multiplicand *= multiplicator
	}
	return multiplicand
}

type Node struct {
	Value string
	Left  *Node
	Right *Node
}
