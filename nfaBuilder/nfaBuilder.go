package nfaBuilder

import (
	"fmt"

)
/* References:
(1) Thompsons algorithim in C: https://swtch.com/~rsc/regexp/regexp1.html
(2) Ian Mcloughlin's Videos (gmit.Learnonline)
*/



/* State struct contains three parts:
(1) Rune representing the character
(2) One edge pointing to another state (character)
(3) Another edge pointing to another state and edge
*/
type state struct {
	symbol rune
	edge1 *state
	edge2 *state
}
// All non deterministic finite automata must have an initial and accept state
// First step in creating nfa draw arrow from initial to accept with E label
type nfa struct {
	initial *state
	accept *state
}

// Takes a postfix regular expression and converts it to an nfa,
// returning a pointer to it
func poregtonfa(pofix string) *nfa{
	nfastack := []*nfa{}

	//loop through expression rune at a time
	for _, r := range pofix{

		switch r {

		// The . operator indicates concatination	
		case '.':
			//pop 2 things off nfa stack
			// LIFO order
			frag2 := nfastack[len(nfastack) -1]
			//remove popped item from the stack
			nfastack = nfastack[:len(nfastack)-1]
			frag1 :=  nfastack[len(nfastack) -1]
			nfastack = nfastack[:len(nfastack)-1]

			//concat : join poppeded items like daisy chain
			frag1.accept.edge1 = frag2.initial
			//pointer
			//push new fragment(joined frags) onto nfa stack
			nfastack = append(nfastack, &nfa{initial: frag1.initial, accept : frag2.accept})
		// The | operator indicates alternation	
		case '|':
			//get last r off stack 
			// LIFO order
			frag2 := nfastack[len(nfastack) -1]
			//remove popped item
			nfastack = nfastack[:len(nfastack)-1]
			frag1 :=  nfastack[len(nfastack) -1]
			nfastack = nfastack[:len(nfastack)-1]

			//create two new states 
			//concat : join poppeded items like daisy chain
			accept :=state{}
			initial := state {edge1: frag1.initial, edge2 : frag2.initial}
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept
			//push new fragment(joined frags) onto nfa stack
			nfastack = append(nfastack, &nfa{initial: &initial, accept :&accept})
		// The Kleane star indicates zero or more	
		case '*':
			frag :=  nfastack[len(nfastack) -1]
			//remove popped item
			nfastack =  nfastack[:len(nfastack)-1]
			//point at oold accpet state & new initial state
			accept := state{}
			//new initial state , point to old initial & new accept
			initial := state{edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = frag.initial //old frag with 2 extra states
			frag.accept.edge2 = &accept

			nfastack = append(nfastack, &nfa{initial: &initial, accept :&accept})
		// The + operator indicates one or more
		case '+':
			//pop one item
			frag :=  nfastack[len(nfastack) -1]
			//remove popped item from stack
			nfastack =  nfastack[:len(nfastack)-1]
			//accept state & new initial state
			accept := state{}
			//one edge going back to itself & another to the accept
			initial := state{edge1: frag.initial, edge2: &accept}

			frag.accept.edge1 = &initial

			nfastack = append(nfastack, &nfa{initial: frag.initial, accept: &accept})
		// The ? operator indicates zero or one
		case '?':
			//pop one item
			frag :=  nfastack[len(nfastack) -1]
			//remove popped item from stack
			nfastack =  nfastack[:len(nfastack)-1]
			//state pointing to popped item and accept state
			initial := state{edge1: frag.initial, edge2: frag.accept}
			// add the nfa to the stack
			nfastack = append(nfastack, &nfa{initial: &initial, accept: frag.accept})
		
		default:
		//	fmt.Print("In default state       ")
		//	fmt.Println(r, string(r))
			accept := state{}
			initial := state{symbol : r, edge1:  &accept }

			nfastack= append (nfastack, &nfa {initial: &initial, accept: &accept})
		}
	}
	
	//should be only one item at end --> nfa you want to return
	//TODO: error checking 
	if len(nfastack) != 1{
		fmt.Println("whoops! ", len(nfastack), nfastack)
		
	}
	return nfastack[0] 
}
//this method runs a given sring through the NFA created returning  a bool indicationg if it is a match/not
func PoMatch (po string, s string) bool {
	
	ismatch := false
	//create automata
	ponfa := poregtonfa(po)
	fmt.Println(ponfa)

	//initial state & everything you can get to from current state
	current := []*state{}
	next := []*state{}


	current = addState(current[:], ponfa.initial, ponfa.accept)

	/* 	Read a char at a time & take all current 
	 	states checking if labelled
	*/
	for _, r := range s {
		//current state
		for _, c := range current {
			//same symbol
			if c.symbol == r{
				next = addState(next[:], c.edge1, ponfa.accept)
			}
		}
		current, next = next, []*state{}
	}

	// loop through current state array and check if there are any accept states
	for _, c := range current {
		//same symbol
		if c == ponfa.accept{
			ismatch = true
			break;
		}
	}

	return ismatch
}
//helper function
func addState(l []*state, s *state, a *state) []*state {
	l = append(l, s)
	//E arrows
	if s != a && s.symbol == 0{  // special value rune
		l = addState(l, s.edge1, a)
		if s.edge2 != nil {
			l = addState(l, s.edge2, a)
		}
	}
	return l
}


