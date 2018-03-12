package main

import (
	"fmt"
)

type state struct {
	symbol rune
	edge1 *state
	edge2 *state
}

type nfa struct {
	nfastack := []*nfa{}
	for _, r:= pxfix{

	}
	
	return nfastack[0]
}

func poregtonfa(pofix string) *nfa{

}
func main(){
	nfa := poregtonfa("ab.c|")
	fmt.Println(nfa)
}