package main

import (
	"fmt"
)

func intopost (infix string) string {
	// using runes instead of string --> UTF-8
	// array of runes (chars)

	// ref for special chars & precendence
	special := map[rune] int{'*':10, '.':9, '|': 8}
	
	postfix := []rune{}
	//store operators from infix
	stack :=[]rune{}





	//cast runes to string
	return string(postfix)
}

func main(){
	//Answer : ab.c*
	fmt.Println("Infix:", "a.b.c*")
	fmt.Println("Postfix: ", intopost("a.b.c*"))

}