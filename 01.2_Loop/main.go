package main

import "fmt"

func main() {

	for i := 100; i <= 200; i++ {

		fmt.Printf("%d \t %b \t %#X \t %T \t %c \t %q \t %#U \n", i, i, i, i, i, i, i)
	}

}
