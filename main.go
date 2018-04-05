/*
--------------------------------------------------------------
Reads in user input via the command line. Parses input &
converts infix notation to postfix if required using the
infixtoPostix package. Builds the NFA from the postfix 
expression using the NFABuilder Package. Reads in user input
to test against the NFA created.
--------------------------------------------------------------
Resources:
	(1): https://stackoverflow.com/questions/20895552/how-to-read-input-from-console-line
*/
package main

import (
	"fmt"
	infixToPostfix "./infixToPostfix"
	nfaBuilder "./NfaBuilder"
)
// Uses postfix notation to geenerate NFA
// Reads in userinput to test regex matching with NFA
func generateNFA(postfix string) {
		//promt the user to enter a string to test against NFA
		fmt.Println("Enter string to test against NFA: ")
		//Read in the user's input from the cmd
		var userTest string
		fmt.Scanln(&userTest)
		//Generaete NFA and test with user's input & print results
		fmt.Println(nfaBuilder.PoMatch(postfix, userTest))
}

func main(){
		// userinput / infix expression / posfix expression
		var userInput, infix, postfix string
		//print a menu to cmd containing user's options
		printMenu()
		//Read in the user's input from the cmd
		fmt.Scanln(&userInput)
		//Loop for user input
		for userInput!="3"{
			// User options
			switch userInput {
				case "1":
					//prompt the user to enter the infix string
					fmt.Println("Enter infix expression to generate NFA: ")
					fmt.Scanln(&infix)
					//Convert string to postfix notation
					postfix := infixToPostfix.IntoPost(infix)
					//create & test NFA
					generateNFA(postfix)
				case "2": 
					fmt.Println("Enter postfix expression to generate NFA: ")
					//prompt the user to enter the postfix string
					fmt.Scanln(&postfix)
					//create & test NFA
					generateNFA(postfix)
				default:
					fmt.Println("Please enter a valid option")
			}
			printMenu()
			fmt.Scanln(&userInput)
		}//for
	}
	
//Prints a selection of options to the user via the cmd
func printMenu(){
	fmt.Println("Enter (1) to generate an NFA from infix notation" +
	"\nEnter (2) to generate an NFA from postfix notation" +
	"\nEnter (3) to Exit")
}
