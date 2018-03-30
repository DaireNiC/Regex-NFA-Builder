package main

import (
	"fmt"
	"os"
	"bufio"
	infixToPostfix "./infixToPostfix"
	nfaBuilder "./NfaBuilder"
	

)

func generateNFA(postfix string) {
		//promt the user to enter a string to test against NFA
		fmt.Print("Enter string to test against NFA: ")
		//Read in the user's input from the cmd
		var scanner  = bufio.NewScanner(os.Stdin)
		scanner.Scan()
		userTest := scanner.Text()
		//Generaete NFA and test with user's input & print results
		fmt.Println(nfaBuilder.PoMatch(postfix, userTest))
}

func main(){
	
		//Loop for user input
		for{
	
			fmt.Println("Enter (1) to generate an NFA from infix notation" +
			"\nEnter (2) to generate an NFA from postfix notation" +
			"\nEnter (3) to Exit\n")
			
			//Read in the user's input from the cmd
			var scanner  = bufio.NewScanner(os.Stdin)
			scanner.Scan()
			userInput := scanner.Text()
	
			// User options
			switch userInput {
				case "1":
					fmt.Println("Enter infix expression to generate NFA: ")
					//promt the user to enter the infix string
					fmt.Print("Enter infix expression: ")
					infixString := scanner.Text()
		
					//Convert string to postfix notation
					postfix := infixToPostfix.IntoPost(infixString)
					//create & test NFA
					generateNFA(postfix)
				case "2": 
					fmt.Println("Enter postfix expression to generate NFA: ")
					//promt the user to enter the postfix string
					fmt.Print("Enter infix expression: ")
					postfix := scanner.Text()
					//create & test NFA
					generateNFA(postfix)
				case "3": 
					fmt.Print("Exiting ....")
					break
				default:
					fmt.Println("Enter one of the above options")
			}
		}
	}
	

		/*
	//==== Example Cases ====//
	postfix := infixToPostfix.IntoPost("abc..")
	//output should be : "ab.c*"
	fmt.Println(postfix)
	// Evalutation is true given input
	fmt.Println(nfaBuilder.PoMatch(postfix, "ab")) //return true
	*/
