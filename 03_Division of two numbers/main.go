package main

import "fmt"

func main() {

	var a, b, c, d int
	fmt.Print("Please Enter the Dividend:  ")
	fmt.Scan(&a)
	fmt.Print("Please Enter the Divisor:  ")
	fmt.Scan(&b)
	c = a / b
	d = a % b
	fmt.Println("The Quotient is: ", c)
	fmt.Println("The Remainder is: ", d)

}
