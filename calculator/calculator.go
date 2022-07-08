package main

import (
	"fmt"
	"os"
	"strconv"
)

func IsValidOperator(op string) bool {
	switch op {
	case
		"*",
		"x",
		"/",
		":",
		"+",
		"-":
		return true
	}
	return false
}

func main() {
	if len(os.Args) > 1 {
		ops := map[string]func(int, int) int{
			"+": func(a, b int) int { return a + b },
			"-": func(a, b int) int { return a - b },
			"*": func(a, b int) int { return a * b },
			"x": func(a, b int) int { return a * b },
			"/": func(a, b int) int { return a / b },
			":": func(a, b int) int { return a / b },
		}

		var x, y int
		var op string
		if len(os.Args) <= 4 {
			op = os.Args[2]
			if !(IsValidOperator(op)) {
				fmt.Println("Invalid Operant")
				return
			}
			x, _ = strconv.Atoi(os.Args[1])
			y, _ = strconv.Atoi(os.Args[3])
		} else {
			op = "*"
			x, _ = strconv.Atoi(os.Args[1])
			y, _ = strconv.Atoi(os.Args[len(os.Args)-1])
		}

		if (x > int(1000000)) || (y > int(1000000)) {
			fmt.Println("Maksimum limit is 1.000.000")
			return
		}
		fmt.Println(ops[op](x, y))
		return
	}
	fmt.Print("Arguments needed to run!")
	return
}
