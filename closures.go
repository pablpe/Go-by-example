package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func stringClosures() func() string {
	str := "hola "
	return func() string {
		str = str + str
		return str
	}
}

func main() {

	nextInt := intSeq()

	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3

	newInts := intSeq()
	fmt.Println(newInts()) // 1

	nextStr := stringClosures()
	fmt.Println(nextStr()) // hola hola
	fmt.Println(nextStr()) // hola hola hola hola
	fmt.Println(nextStr()) //  ola hola hola hola hola hola hola hola
}
