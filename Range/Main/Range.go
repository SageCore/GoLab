package main

import "fmt"

func main() {
	//Here we use range to sum the numbers in a slice. Arrays work like this too.
	nums := []int{54, 45, 12, 64}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	minus := 0
	for _, num := range nums {
		minus -= num
	}
	fmt.Println("minus:", minus)

	multiplication := 0
	for _, num := range nums {
		multiplication *= num
	}
	fmt.Println("multiplication:", multiplication)

	division := 0
	for _, num := range nums {
		division /= num
	}
	fmt.Println("division:", division)
	//range on arrays and slices provides both the index and value for each entry. Above we didn’t need the index, so we ignored it with the blank identifier _. Sometimes we actually want the indexes though.
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	//range on map iterates over key/value pairs.
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	//range can also iterate over just the keys of a map.
	for k := range kvs {
		fmt.Println("key:", k)
	}
	//range on strings iterates over Unicode code points. The first value is the starting byte index of the rune and the second the rune itself.
	for i, c := range "fmt" {
		fmt.Println(i, c)
	}
}
