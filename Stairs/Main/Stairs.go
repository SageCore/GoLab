package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("stairs:", s)

	s[0] = "__\n"

	for i := 0; i < 50; i++ {
		s = append(s, "__\n")
	}

	fmt.Println("stairs:", s)


}
