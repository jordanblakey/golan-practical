package main

import "fmt"

// func add(x float64, y float64) float64 {
func add(x, y float32 /* arg type */) float32 /* return type */ { // Shorthand for typing
	return x + y
}

func multiple(a, b string) (string, string /* if returning multiple values, must specify return type for each */) {
	return a, b
}

func main() {
	// var num1 float64 = 5.6
	// var num2 float64 = 9.5
	// var num1, num2 float64 = 5.6, 9.5 // shorthand assignment
	// const x int = 5 // Constant declaration
	num1, num2 := 5.6, 9.5 // Use type inferencing each time the variable is called
	w1, w2 := "Hey", "there"

	// You only need to specify a type when declaring but not assigning a variable,
	// or if you want the type to be different than what's inferred.
	// Otherwise, the variable's type will be the same as that of the right-hand side of the assignment.
	var a float32 = 62
	var b = float64(a)
	x := a // x will be type int

	fmt.Println(b)

	fmt.Println(add(float32(num1), float32(num2)))
	fmt.Println(multiple(w1, w2))
}
