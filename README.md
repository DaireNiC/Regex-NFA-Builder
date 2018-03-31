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

  
## How it works

    1. Reads in user input in postfix/infix notation
    2. If input is given in infix, parses the regular expression from infix to postfix notation.
    3. Builds a series of small NFA’s representing parts of the regular expression.
    4. Uses the smaller NFA’s to create the overall NFA.
    5. User enters input to implememnt the matching algorithm using the newly created NFA.
    
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

### Parsing regex from infix to postfix 
Resources:
* (1): http://interactivepython.org/runestone/static/pythonds/BasicDS/InfixPrefixandPostfixExpressions.html
* (2): https://www.geeksforgeeks.org/stack-set-2-infix-to-postfix/
* (3): [Youtube Video on Shunting Yard Algorithm](https://www.youtube.com/watch?v=HJOnJU77EUs)
* (4): [Go Lang video provided by Ian McLoughlin](https://web.microsoftstream.com/video/9d83a3f3-bc4f-4bda-95cc-b21c8e67675e) 
