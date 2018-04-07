/* A series of tests for the postfix notation to NFA builder
Guidance on test structure & layout in GoLang from :
(1) https://blog.golang.org/subtests
(2) https://medium.com/@matryer/5-simple-tips-and-tricks-for-writing-unit-tests-in-golang-619653f90742
*/
package nfaBuilder

import (
	"testing"
)

//Table structure for testdata
var nfabuilderTests = []struct {
	postfixRegex string // regex string given in post ix notation
	n       string //input
	expected         bool //expected result
  }{
	// | operator --> (OR)
	{"ab|", "a", true},
	{"ab|", "b", true},
	{"ab|", "ab", false},
	// . operator --> concatenation (AND)
	{"ab.", "ab", true},
	{"ab.", "b", false},
	{"ab.", "a", false},
	// + operator --> one or more
	{"10.1.+", "", false},
	{"10.1.+", "101101101", true},
	{"10.1.+", "101", true},
	// ? operator --> zero or one
	{"10.1?|", "", true},
	{"10.1?|", "1", true},
	{"10.1?|", "1111", false},
	// * Kleene star --> zero or more	
	{"10.1*|", "", true},
	{"10.1*|", "1", true},
	{"10.1*.", "10111", true},
	{"10.1*.", "101", true},
	{"10.1*|", "0", false},
	// Combined Expressions Tests
	{"ab.f*|", "ab", true},
	{"ab.f*|", "ffff", true},
	{"eP.9*.", "eP999", true},
	{"eP.9*|", "EPPPPPPPPPPPPPPpp", false},

  }

/*	Tests the NFA's built from the given postFix regex
 	with input string (n) & asserts true/false
 */
 func TestNFABuilder(t *testing.T) {
	for _, tt := range nfabuilderTests {
		actual := PoMatch(tt.postfixRegex, tt.n)
		if actual != tt.expected {
		  t.Errorf("\nPostfix Regex: %q\nInput String: %q\nExpected: %b, \nActual %b", tt.postfixRegex, tt.n, tt.expected, actual)
		}
	  }
	}
