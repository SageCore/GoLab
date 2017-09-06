//Run code
package main

import "fmt"

//Here’s a function that takes two ints and returns their sum as an int.
func plus(a int, b int) int {
	//Go requires explicit returns, i.e. it won’t automatically return the value of the last expression.
	return a + b
}

func minus(a int, b int) int {
	//Go requires explicit returns, i.e. it won’t automatically return the value of the last expression.
	return a - b
}

func multiply(a int, b int) int {
	//Go requires explicit returns, i.e. it won’t automatically return the value of the last expression.
	return a * b
}

func divide(a int, b int) int {
	//Go requires explicit returns, i.e. it won’t automatically return the value of the last expression.
	return a / b
}

//When you have multiple consecutive parameters of the same type, you may omit the type name for the like-typed parameters up to the final parameter that declares the type.
func plusPlus(a, b, c int) int {
	return a + b + c
}
func main() {
	//Call a function just as you’d expect, with name(args).
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	ser := minus(54, 32)
	fmt.Println("54 - 32 =", ser)

	her := multiply(34, 98)
	fmt.Println("34 * 98 =", her)

	ker := divide(81, 9)
	fmt.Println("81 / 9 =", ker)
}
