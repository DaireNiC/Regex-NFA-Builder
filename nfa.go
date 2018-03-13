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
	initial *state
	accept *state
}

func poregtonfa(pofix string) *nfa{
	nfastack := []*nfa{}

	//loop through exp rune at a time
	for _, r := range pofix{

		switch r{
		//pop 2 things off nfa stack
		case '.':
			//get last r off stack 
			// LIFO order
			frag2 := nfastack[len(nfastack) -1]
			//remove popped item
			nfastack = nfastack[:len(nfastack)-1]
			frag1 :=  nfastack[len(nfastack) -1]
			nfastack = nfastack[:len(nfastack)-1]

			//concat : join poppeded items like daisy chain
			frag1.accept.edge1 = frag2.initial
			//pointer
			//push new fragment(joined frags) onto nfa stack
			nfastack = append(nfastack, &nfa{initial: frag1.initial, accept : frag2.accept})

		case '|':
			//get last r off stack 
			// LIFO order
			frag2 := nfastack[len(nfastack) -1]
			//remove popped item
			nfastack = nfastack[:len(nfastack)-1]
			frag1 :=  nfastack[len(nfastack) -1]
			nfastack = nfastack[:len(nfastack)-1]

			//create two new states 
			accept :=state{}
			initial := state {edge1: frag1.initial, edge2 : frag2.initial}
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			//concat : join poppeded items like daisy chain
			frag1.accept.edge1 = frag2.initial
			//pointer
			//push new fragment(joined frags) onto nfa stack
			nfastack = append(nfastack, &nfa{initial: &initial, accept :&accept})

		case '*':
			frag :=  nfastack[len(nfastack) -1]
			//remove popped item
			nfastack =  nfastack[:len(nfastack)-1]
			//point at oold accpet state & new initial state
			accept := state{}
			//new initial state , point to old initial & new accept
			initial := state{edge1: frag.initial, edge2: frag.accept}
			frag.accept.edge1 = frag.initial //old frag with 2 extra states
			frag.accept.edge2 = &accept

			nfastack = append(nfastack, &nfa{initial: &initial, accept :&accept})
			
		default:
			accept := state{}
			initial := state{symbol : r, edge1:  &accept }

			nfastack= append (nfastack, &nfa {initial: &initial, accept: &accept})
		}
	}
	
	//should be only one item at end --> nfa you want to return
	//TODO: error checking 
	return nfastack[0] 
}

func main(){
	nfa := poregtonfa("ab.c|")
	fmt.Println(nfa)
}