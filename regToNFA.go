package main

import (
	"fmt"
	//postfixMaker "./utils"
	NfaBuilder "./NfaBuilder"
	

)
func main(){
	//test1
//	fmt.Println(NfaBuilder.PoMatch("a|c*b", "ab"))
	fmt.Println(NfaBuilder.PoMatch("ab.c?|", "c")) //return true
//	fmt.Println(NfaBuilder.PoMatch("a+c*|", "ac"))
//	fmt.Println(NfaBuilder.PoMatch(postfixMaker.IntoPost("a?b?c?"), "abc"))
//	fmt.Println(NfaBuilder.PoMatch("ab.c+", "ac"))
//	fmt.Println(NfaBuilder.PoMatch("ab?c*", "abcc"))
//	fmt.Println(NfaBuilder.PoMatch("ab?c*", "abbcc"))
	
}