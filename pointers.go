package main

import "fmt"

func printRefValue(val int) {
	fmt.Print(val)
}
func printPtrValue(val *int) {
	fmt.Print(*val)
}
func main() {
	num := 3
	printRefValue(num)
	printPtrValue(&num)
}
