package main

import "fmt"

func main() {

	x := 0
	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())

}
// Normally you cannot declare a function witin a function you have to adopt a different way.
//  Here i declared a function increment within the main function.This is call closure there are many different ways of doing it.

 /*
Defination of Closure
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope

anonymous function
a function without a name

func expression
assigning a func to a variable
*/
