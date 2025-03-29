package main

import "fmt"

func finalPositionOfSnake(n int, commands []string) int {
	i := 0
	j := 0
	for k := 0; k < len(commands); k++ {
		switch commands[k] {
		case "UP":
			i -= 1
		case "DOWN":
			i += 1
		case "LEFT":
			j -= 1
		case "RIGHT":
			j += 1
		}
	}
	result := (i * n) + j
	return result
}

func main() {
	n := 3
	commands := []string{"DOWN", "RIGHT", "UP"}
	fmt.Print(finalPositionOfSnake(n, commands))
}
