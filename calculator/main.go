package main

import "fmt"

func main() {
	var num1, num2 float64
	var operation string

	fmt.Println("===КАЛЬКУЛЯТОР SLABENK1Y===")
	fmt.Println()

	fmt.Print("Первое число: ")
	fmt.Scan(&num1)

	fmt.Print("Операция (+ - * /):")
    fmt.Scan(&operation)

	fmt.Print("Второе число: ")
	fmt.Scan(&num2)

	fmt.Println()

	if operation == "+" {
		fmt.Println("Результат:", num1, "+", num2, "=", num1+num2)
	} else if operation == "-" {
		fmt.Println("Результат:", num1, "-", num2, "=", num1-num2)
	} else if operation == "*" {
		fmt.Println("Результат:", num1, "*", num2, "=", num1*num2)
	} else if operation == "/" {
		if num2 != 0 {
			fmt.Println("Результат:", num1, "/", num2, "=", num1/num2)
		} else {
			fmt.Println("Ошибка: деление на ноль")
		}
	} else {
		fmt.Println("Ошибка: неизвестная операция")
	}
	}