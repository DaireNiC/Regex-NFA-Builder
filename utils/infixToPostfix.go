package infixtoPostfix

import (
	"fmt"
)

func intopost (infix string) string {
	// using runes instead of chars, in GO --> UTF-8

	// ref for special chars & precendence
	specials := map[rune] int{'*':10, '.':9, '|': 8}
	// array of runes (chars)
	postfix := []rune{}
	//store operators from infix
	stack :=[]rune{}

	//loop through infix, return index of char currently reading
	//r is the char*(rune) at that index
	for _, r := range infix {
		switch {
			case r == '(':
				//put bracket on stack
				stack = append(stack,r)
			case r == ')':
				//while last char on stack not the closing bracket just read
				for stack[len(stack) -1] != '(' {
					postfix = append(postfix, stack[len(stack)-1])
					stack = stack[:len(stack) -1] //everything in stack up to closing bracket
				}
				stack = stack[:len(stack) -1]
			// if element isn't in map, return 0
			case specials[r] > 0:
				for len(stack) > 0 && specials[r] <= specials[stack[len(stack) -1]]{
					postfix = append(postfix, stack[len(stack)-1])
					stack = stack[:len(stack) -1]
				}
				stack = append(stack,r)
			default:
				postfix = append(postfix, r)
		}
	}

	for len(stack) > 0{
		//take top el of stack & put on top of postfixs
		postfix = append(postfix, stack[len(stack)-1])
		stack = stack[:len(stack) -1]
	}

	//cast runes to string
	return string(postfix)
}

func main(){
	//Answer : ab.c*
	fmt.Println("Example 1: ")
	fmt.Println("Infix:", "a.b.c*")
	fmt.Println("Postfix: ", intopost("a.b.c*"))

	//Answer : ab*c.b|
	fmt.Println("\nExample 2:")
	fmt.Println("Infix: ", "a*b.c|b")
	fmt.Println("Postfix: ", intopost("a*b.c|b"))

	//Answer : a|c*b
	fmt.Println("Infix: ", "a|c*b")
	fmt.Println("Postfix: ", intopost("a|c*b"))

}