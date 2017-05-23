package main

import "fmt"

func main() {

	var Name string
	fmt.Print("Please Enter Your Name: \t")
	fmt.Scan(&Name) 						// only works if there is no space between your name 
	fmt.Println("Hello Mr.", Name)

}

