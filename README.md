# Regex-creator

## Given Problem Statement
You must write a program in the [Go Programming Language](http://golang.org) that can
build a non-deterministic finite automaton (NFA) from a regular expression,
and can use the NFA to check if the regular expression matches any given
string of text. You must write the program from scratch and cannot use the
regexp package from the Go standard library nor any other external library.
A regular expression is a string containing a series of characters, some
of which may have a special meaning. For example, the three characters
“.”, “|”, and “∗
” have the special meanings “concatenate”, “or”, and “Kleene
star” respectively. So, 0.1 means a 0 followed by a 1, 0|1 means a 0 or a 1,
and 1∗ means any number of 1’s. These special characters must be used in
your submission. Other special characters you might consider allowing as input are brackets
“()” which can be used for grouping, “+” which means “at least one of”, and
“?” which means “zero or one of”. You might also decide to remove the
concatenation character, so that 1.0 becomes 10, with the concatenation
implicit.
You may initially restrict the non-special characters your program works
with to 0 and 1, if you wish. However, you should at least attempt to expand
these to all of the digits, and the characters a to z, and A to Z.

## Simplified 
  * Go program to build NFA from regex
  * Use NFA to check if regex matches string
  * Regex => regular characters + special characters 
  
### The program can be broken into the following steps:
    1. Parse the regular expression from infix to postfix notation.
    2. Build a series of small NFA’s for parts of the regular expression.
    3. Use the smaller NFA’s to create the overall NFA.
    4. Implement the matching algorithm using the NFA.
    
### Parsing regex from infix to postfix 
Resources:
* (1): http://interactivepython.org/runestone/static/pythonds/BasicDS/InfixPrefixandPostfixExpressions.html
* (2): https://www.geeksforgeeks.org/stack-set-2-infix-to-postfix/
