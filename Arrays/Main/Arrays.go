//Run code
package main

import "fmt"

func main() {
	//Here we create an array a that will hold exactly 5 ints. The type of elements and length are both part of the array’s type. By default an array is zero-valued, which for ints means 0s.
	var a [5]int
	fmt.Println("emp:", a)
	//We can set a value at an index using the array[index] = value syntax, and get a value with array[index].
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	//The builtin len returns the length of an array.
	fmt.Println("len:", len(a))
	//Use this syntax to declare and initialize an array in one line.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	c := [5]string{"get", "me", "a", "book", "Please!"}
	fmt.Println("", c[0])
	fmt.Println("", c[1])
	fmt.Println("", c[2])
	fmt.Println("", c[3])
	fmt.Println("", c[4])

	fmt.Println("\n", c[4])
	fmt.Println("", c[0])
	fmt.Println("", c[1])
	fmt.Println("", c[2])
	fmt.Println("", c[3])
	//Array types are one-dimensional, but you can compose types to build multi-dimensional data structures.
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}