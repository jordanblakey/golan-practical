package main

import "fmt"

func main() {
	x := 15
	a := &x // & reads the memory address of variable x and stores it in a

	fmt.Println(a)  // value of a, which represents a memory address
	fmt.Println(*a) // the value at the memory address stored in a
	*a = 5          // reassigns the value of the memory addres in a
	fmt.Println(x)  // the value of the original var x, reassigned in the prev line
	*a = *a * *a    // reassign the value at the memory address stored in a, times itself
	fmt.Println(x)  // print the original value
	fmt.Println(*a) // print the original value
}
