package main

import (
	"fmt"
	infixToPostfix "./InfixToPostfix"
	nfaBuilder "./NfaBuilder"
	

)
func main(){
	//==== Test Cases ====//
	postfix := infixToPostfix.IntoPost("(a.b)|c")
	//output should be : "ab.c*"
	fmt.Println(postfix)
	// Evalutation is true given input
	fmt.Println(nfaBuilder.PoMatch(postfix, "ab")) //return true
}