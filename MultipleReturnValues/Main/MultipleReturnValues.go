//Run code
package main

import "fmt"

//The (int, int) in this function signature shows that the function returns 2 ints.
func vals() (int, int, int) {
	return 32, 2, 7
}
func main() {
	//Here we use the 2 different return values from the call with multiple assignment.
	a, b, c := vals()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//If you only want a subset of the returned values, use the blank identifier _.
	_, _, d := vals()
	fmt.Println(d)

	//Here is an other example

	m, l, e := vals()
	fmt.Println(e, l, m)
	fmt.Println(e)

	_, _, k := vals()
	fmt.Println(k)
}
