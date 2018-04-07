/* 
---------------------------------------------------------------
A series of tests for the infix to postfix notation converter
---------------------------------------------------------------
Guidance on test structure & layout in GoLang from :
(1) https://blog.golang.org/subtests
(2) https://medium.com/@matryer/5-simple-tips-and-tricks-for-writing-unit-tests-in-golang-619653f90742
*/
package infixToPostfix

import (
	"testing"

)

//Table structure for testdata
var infixToPostfixTests = []struct {
	n string // infix regex string
	expected         string //expected result
  }{
//==|| INFIX |  POSTFIX ||==//  
	{"(a.b)*c", "ab.c*"},
	{"(a.(b|d))*", "abd|.*"},
	{"a.(b.b)+.c", "abb..c.+"},
  }

/*	Tests the conversion from the given infix regex string (n)
  	to postfix notation
 */
 func TestInfixTosPostfix(t *testing.T) {
	for _, tt := range infixToPostfixTests {
		actual := IntoPost(tt.n)
		if actual != tt.expected {
		  t.Errorf("\nInfix Regex: %q \nExpected Postfix: %q, \nActual %q", tt.n,tt.expected, actual)
		}
	  }
	}
