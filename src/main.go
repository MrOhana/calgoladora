package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	for {
		userInput := GetUserInputOption()

		if !ValidateUserInput(userInput) {
			fmt.Println("Invalid input. Choose a valid option.")
			continue
		}

		if userInput == "5" {
			fmt.Println("Thanks to use my calGOlator. See you soon!")
			break
		}

		num1 := GetUserInputNumber("Choose the first number...")
		num2 := GetUserInputNumber("Choose the second number...")

		resultOperation := PerformOperation(userInput, num1, num2)
		fmt.Printf("The result of the operation is: %.2f\n", resultOperation)
	}
}

func GetUserInputOption() string {
	fmt.Println("################################################")
	fmt.Println("Please choose one of the following options:")
	fmt.Println("1. Soma")
	fmt.Println("2. Subtração")
	fmt.Println("3. Multiplicação")
	fmt.Println("4. Divisão")
	fmt.Println("__________________")
	fmt.Println("5. Exit")
	fmt.Println("################################################")

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter input: ")
	scanner.Scan()
	return scanner.Text()
}

func GetUserInputNumber(message string) float64 {
	inputNumber := GetUserInput(message)
	resNumber, err := strconv.ParseFloat(inputNumber, 64)
	if err != nil {
		log.Fatalf("Invalid input. Please enter a valid number. Error: %v", err)
	}

	return resNumber
}

func GetUserInput(message string) string {
	fmt.Println(message)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(">> ")
	scanner.Scan()
	return scanner.Text()
}

func ValidateUserInput(input string) bool {
	return (len(input) > 0 &&
		(input == "1" || input == "2" || input == "3" || input == "4" || input == "5"))
}

func PerformOperation(operation string, num1 float64, num2 float64) float64 {
	var result float64

	switch operation {
	case "1":
		result = num1 + num2
	case "2":
		result = num1 - num2
	case "3":
		result = num1 * num2
	case "4":
		if num2 == 0 {
			fmt.Println("Error: Division by zero is not allowed.")
			return 0.0
		}
		result = num1 / num2
	default:
		fmt.Println("Invalid operation.")
	}

	return result
}
