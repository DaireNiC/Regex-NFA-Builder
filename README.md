

  
# Regex to NFA Builder
A command-line application that converts regular expressions to NFAs. The application then allows users to test input against the built NFA. 
## About
Name: Daire Ní Chatháin
ID: G00334757

This application was created as part of my Graph Theory Module in GMIT. The code here is adapted from many sources, all referenced within the source code or in this readMe file. 

## Given Problem Statement

> You must write a program in the [Go Programming
> Language](http://golang.org) that can build a non-deterministic finite
> automaton (NFA) from a regular expression, and can use the NFA to
> check if the regular expression matches any given string of text. You
> must write the program from scratch and cannot use the regexp package
> from the Go standard library nor any other external library. A regular
> expression is a string containing a series of characters, some of
> which may have a special meaning. For example, the three characters
> “.”, “|”, and “∗ ” have the special meanings “concatenate”, “or”, and
> “Kleene star” respectively. So, 0.1 means a 0 followed by a 1, 0|1
> means a 0 or a 1, and 1∗ means any number of 1’s. These special
> characters must be used in your submission. Other special characters
> you might consider allowing as input are brackets “()” which can be
> used for grouping, “+” which means “at least one of”, and “?” which
> means “zero or one of”. You might also decide to remove the
> concatenation character, so that 1.0 becomes 10, with the
> concatenation implicit. You may initially restrict the non-special
> characters your program works with to 0 and 1, if you wish. However,
> you should at least attempt to expand these to all of the digits, and
> the characters a to z, and A to Z.

## Problem Simplified 
  * Go program to build NFA from regex
  * Use NFA to check if regex matches string
  * Regex => regular characters + special characters 
 ![flow_chart](https://github.com/DaireNiC/Regex-NFA-Builder/blob/master/media/flow_chart.JPG)
 
## How it works

    1. Reads in user input in postfix/infix notation
    2. If input is given in infix, parses the regular expression from infix to postfix notation.
    3. Builds a series of small NFA’s representing parts of the regular expression.
    4. Uses the smaller NFA’s to create the overall NFA.
    5. User enters input to implememnt the matching algorithm using the newly created NFA.
   

![cmd](https://github.com/DaireNiC/Regex-NFA-Builder/blob/master/media/cmd_01.JPG)  


## Running the program
[Go](https://golang.org) must be installed to run the code. Follow official guidelines here.

### Clone this repo
```bash
git clone https://github.com/DaireNiC/Regex-NFA-Builder
```
### Navigate to the folder
```bash
cd Regex-NFA-Builder
```
### Compile the Go code
```go
go build main.go
```
### Run the exe:
```bash
./main.exe
```
## Code & Program Design

### Converting Input from Infix to Postfix Notation

The NFA is created from postfix notation. This requires converting the user input from infix to postfix notation.  A stack-like structure is implemented to facilitate this. The code for this conversion is found in the [infixToPostfix.go](https://github.com/DaireNiC/Regex-NFA-Builder/blob/master/InfixToPostfix/infixToPostfix.go) file.  The Shunting Yard Algorithim is used as the basis for this code. 

Resources:
* (1): http://interactivepython.org/runestone/static/pythonds/BasicDS/InfixPrefixandPostfixExpressions.html
* (2): https://www.geeksforgeeks.org/stack-set-2-infix-to-postfix/
* (3): [Youtube Video on Shunting Yard Algorithm](https://www.youtube.com/watch?v=HJOnJU77EUs)
* (4): [Go Lang video provided by Ian McLoughlin](https://web.microsoftstream.com/video/9d83a3f3-bc4f-4bda-95cc-b21c8e67675e) 

### Building the NFA
Thompsons's Construction Algorithim is used to build NFA's from the user's input in postfix notation. Both NFA's and Regular expressions have the same computational power, therefore any language that can be recognised by a regex can also be represented as a NFA, and visa versa.  (link)

Thompson's Construction uses fragments to create an NFA that recognises a given regular expression in postfix notation. A fragment in this case, is a smaller NFA used to build a larger, overall NFA. 

A stack structure is used to store the fragments for the creation of the overall NFA. In Thompsons Construction, all fragments have a single initial state(required for an NFA) and a single final state. I used the following NFA struct  to represent this:
```go
type  nfa  struct {
	initial *state
	accept *state
}
```

 Normal characters (*see below for special characters*) are pushed to the stack in the from of a fragment as detailed below.  As it is an NFA, the empty regular expression **ε** is also included.
![non_special_character](https://github.com/DaireNiC/Regex-NFA-Builder/blob/master/media/nonspecial_char.JPG)

The struct used to represent the states and arrows is shown below. In Thompson's Construction, every state has two arrows at max.  In the case of only one arrow from a state, the second edge will simply be ignored. 
```go
type  state  struct {
	symbol rune // if ε, will be represented as a 0 value
	edge1 *state //pointers to other states (similar to linked list)
	edge2 *state 
}
```
 The following are considered special characters. These characters pop from & push  elements to the stack. 
 
 ||Special Characters Recognised|		|
|--|--|--|
| &#124; | OR | ![or](https://github.com/DaireNiC/Regex-NFA-Builder/blob/master/media/OR.JPG)|
| .| concatenation |![concat](https://github.com/DaireNiC/Regex-NFA-Builder/blob/master/media/concat.JPG)|
| * | zero or more |![zero_or_more](https://github.com/DaireNiC/Regex-NFA-Builder/blob/master/media/zero_or_more.JPG)|
|  +  | one or more |![enter image description here](https://github.com/DaireNiC/Regex-NFA-Builder/blob/master/media/one_or_more.JPG)|
| ? | zero or one|![enter image description here](https://github.com/DaireNiC/Regex-NFA-Builder/blob/master/media/zero_or_one..JPG)|
 



## Testing
In an effort to both understand and get experience with Unit Testing using GoLang I created a test suite. I used the GoLang Docs [here](https://golang.org/pkg/testing/) for guidance on creating and running tests using Go. 

The first test tests the conversion of infix notation to postfix. The test coverage in all tests is 97~% and covers a range of cases.  

![test_coverage](https://github.com/DaireNiC/Regex-NFA-Builder/blob/master/media/nfabuilder_testcoverage.JPG)

The second tests the postfix to NFA creation. Test tables were used for efficient creation of tests and their assertions. Guidance on the use of test tables was found on [this blog](https://blog.alexellis.io/golang-writing-unit-tests/). 

### Running the Tests

 -  Navigate to the folder
```bash
	cd Regex-NFA-Builder/Tests
```
 - Execute the tests
```go
	go test
```
## Extras relevant to Project

 - The regex engine can both create NFAs & test them against strings over the alaphabet a to z, and A to Z. It can also recognise all digits.
 - The special character "+", which means 'at least one of' was added as an extra feature
  - The special character "?", which means 'zero or one of' was added as an extra feature 
  - Options for infix or postfix input
 - Comprehensive tests were added to the project for robustness
 - An extensive github history indicating the time spent on research, coding, documentation and referencing where due.
 - Consistent work ethic, committing week by week to the repo.


----------
## Research
###  Finite Automata & Regular Expressions
While researching both the theory and implementation of finite automata, I used The Theory of Computation by Michael Sipser along with Dr.Ian McLoughlin's notes (lecturer @ GMIT)  as my primary resources. 

For the creation of the Regex to NFA builder, the paper [Regular Expression Matching Can Be Simple And Fast](https://swtch.com/~rsc/regexp/regexp1.html) by Russ Cox was the primary basis for the code written in this application. Cox provides an implementation in C, which I adapted to run as a Go program.


### Go application Structure, GoDocs, Code Comments
Guidance on structuring code in Go (comments, tests, best practises, documentation) was found on the following blogs & articles :
 - https://medium.com/@benbjohnson/structuring-applications-in-go-3b04be4ff091
 - http://blog.el-chavez.me/2013/08/29/golang-documenting-package-examples/
 - http://whipperstacker.com/2015/10/14/idiomatic-doc-comments-document-your-function-not-your-function-signature/
 - https://blog.golang.org/godoc-documenting-go-code
