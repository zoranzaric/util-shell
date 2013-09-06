package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}
	numbers := []float64{}
	for _, arg := range os.Args[2:] {
		value, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			os.Exit(1)
		}
		numbers = append(numbers, value)
	}
	switch os.Args[1] {
	case ">":
		if len(numbers) != 2 {
			os.Exit(1)
		}
		if numbers[0] > numbers[1] {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	case "<":
		if len(numbers) != 2 {
			os.Exit(1)
		}
		if numbers[0] < numbers[1] {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	case "==":
		if len(numbers) != 2 {
			os.Exit(1)
		}
		if numbers[0] == numbers[1] {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	case ">=":
		if len(numbers) != 2 {
			os.Exit(1)
		}
		if numbers[0] >= numbers[1] {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	case "<=":
		if len(numbers) != 2 {
			os.Exit(1)
		}
		if numbers[0] <= numbers[1] {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	case "+":
		if len(numbers) != 2 {
			os.Exit(1)
		}
		fmt.Println(numbers[0] + numbers[1])
	case "-":
		switch len(numbers) {
		case 1:
			fmt.Println(-numbers[0])
		case 2:
			fmt.Println(numbers[0] - numbers[1])
		default:
			os.Exit(1)
		}
	case "*":
		if len(numbers) != 2 {
			os.Exit(1)
		}
		fmt.Println(numbers[0] * numbers[1])
	case "/":
		if len(numbers) != 2 {
			os.Exit(1)
		}
		fmt.Println(numbers[0] / numbers[1])
	case "e":
		if len(numbers) != 0 {
			os.Exit(1)
		}
		fmt.Println(math.E)
	case "pi":
		if len(numbers) != 0 {
			os.Exit(1)
		}
		fmt.Println(math.Pi)
	case "ceil":
		if len(numbers) != 1 {
			os.Exit(1)
		}
		fmt.Println(math.Ceil(numbers[0]))
	case "floor":
		if len(numbers) != 1 {
			os.Exit(1)
		}
		fmt.Println(math.Ceil(numbers[0]))
	case "log":
		if len(numbers) != 1 {
			os.Exit(1)
		}
		fmt.Println(math.Log(numbers[0]))
	case "pow":
		if len(numbers) != 2 {
			os.Exit(1)
		}
		fmt.Println(math.Pow(numbers[0], numbers[1]))
	case "sqrt":
		if len(numbers) != 1 {
			os.Exit(1)
		}
		fmt.Println(math.Sqrt(numbers[0]))
	default:
		fmt.Println("not supported:", os.Args[1])
		os.Exit(1)
	}
	os.Exit(0)
}
