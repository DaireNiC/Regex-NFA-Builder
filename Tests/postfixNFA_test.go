package main

import (
	"testing"
	"../NfaBuilder"
)

func TestPostFixToNFA(t *testing.T) {
	//==== Test Cases ====//

	// | operator --> (OR)
	fmt.Println("\n==== (|) operator Tests =====")
	var bool  = NfaBuilder.PoMatch("ab|", "a")
	fmt.Println(NfaBuilder.PoMatch("ab|", "a")) //return true
	fmt.Println(NfaBuilder.PoMatch("ab|", "b")) //return true
	fmt.Println(NfaBuilder.PoMatch("ab|", "ab")) //return false

	// . operator --> concatenation (AND)
	fmt.Println("\n==== (.) operator Tests =====")
	fmt.Println(NfaBuilder.PoMatch("ab.", "ab")) //return true
	fmt.Println(NfaBuilder.PoMatch("ab.", "b")) //return false
	fmt.Println(NfaBuilder.PoMatch("ab.", "a")) //return false

	// + operator --> one or more
	fmt.Println("\n==== (+) operator Tests =====")
	fmt.Println(NfaBuilder.PoMatch("10.1+|", "1")) //return true
	fmt.Println(NfaBuilder.PoMatch("10.1+|", "111111")) //return true
	fmt.Println(NfaBuilder.PoMatch("10.1+|", "")) //return false

	// ? operator --> zero or one
	fmt.Println("\n==== (?) operator Tests =====")
	fmt.Println(NfaBuilder.PoMatch("10.1?|", "")) //return true
	fmt.Println(NfaBuilder.PoMatch("10.1?|", "1")) //return true
	fmt.Println(NfaBuilder.PoMatch("10.1?|", "111")) //return 
	

	// * Kleene star --> zero or more	
	fmt.Println("\n==== Kleene Star Tests =====")
	fmt.Println(NfaBuilder.PoMatch("10.1*|", "")) //return true
	fmt.Println(NfaBuilder.PoMatch("10.1+|", "1")) //return true
	fmt.Println(NfaBuilder.PoMatch("10.1+|", "111")) //return true
	fmt.Println(NfaBuilder.PoMatch("10.1+|", "0")) //return false

	// Combined Expressions Tests
	fmt.Println("\n==== Combined Expressions Tests =====")
	fmt.Println(NfaBuilder.PoMatch("ab.f*|", "ffffff")) //return true
	fmt.Println(NfaBuilder.PoMatch("eP.10.Q.q.*|", "")) //return true
}