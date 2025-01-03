package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func main() {
	var result float64
	var operator string

	color.Cyan("Welcome to calculator!")
	scanner := bufio.NewScanner(os.Stdin)

	getNumber := func(prompt string) float64 {
		for {
			fmt.Println(prompt)
			scanner.Scan()
			input := scanner.Text()
			num, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
			if err != nil {
				color.Red("Invalid input! please enter a valid number.")
				continue
			}
			return num
		}
	}

	result = getNumber("Enter first number: ")
	for {

		fmt.Println("operator ('q' to quit): ")
		scanner.Scan()
		operator = strings.TrimSpace(scanner.Text())
		if operator == "q" {
			color.Red("Exiting...")
			break
		}

		num := getNumber("Enter next number: ")

		switch operator {
		case "q":
			color.Red("Exiting...")
			break
		case "+":
			result += num
		case "-":
			result -= num
		case "*":
			result *= num
		case "/":
			if num != 0 {
				result /= num
			} else {
				color.Red("Error: Division by zero!")
				continue
			}
		default:
			color.Red("Error: Invalid operator!")
			continue
		}

		color.Green("Result: %.2f", result)
	}

}
