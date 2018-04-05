/* 
--------------------------------------------------------------------
Implementation of the Shunting Yard Algorithim in Go
Used to convert from Infix notation to Postfix
--------------------------------------------------------------------
 Resources:
 	(1): https://en.wikipedia.org/wiki/Shunting-yard_algorithm
	(2): http://eddmann.com/posts/shunting-yard-implementation-in-java/	
	(3): https://www.youtube.com/watch?v=HJOnJU77EUs
	(4): Video provided by Ian McLoughlin @GMIT
*/
package infixToPostfix

//IntoPost converts and returns a string given in infix to postfix notation
func IntoPost (infix string) string {
	// using runes instead of chars, in GO --> UTF-8
	// ref for special chars stored in map struct
	// guidelines: http://www.boost.org/doc/libs/1_56_0/libs/regex/doc/html/boost_regex/syntax/basic_extended.html#boost_regex.syntax.basic_extended.operator_precedence
	specials := map[rune] int{
		'*':10,
		'.':9,
		'+':8,
		'?':7,
		'|':6,
	 }

	// array of runes (chars)
	postfix := []rune{}
	//store operators from infix
	stack :=[]rune{} //stack for LIFO operations

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
			// if the rune is a special character
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
