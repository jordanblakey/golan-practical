// TO GET DOCUMENTATION ON PACKAGE OR FUNCTION
// $ godoc <package> <function>
// $ godoc fmt Println
// $ godoc math/rand Intn

package main // This will always go at the top level

import (
	"fmt" // Function for
	"math"
	"math/rand"
)

func foo() {
	fmt.Println("This is the foo function.")
}

func random() {
	fmt.Println("A random number from 1-100:", rand.Intn(100))
}

func main() {
	fmt.Println("The square root of 4 is", math.Sqrt(4))
	foo()
	random()
}
