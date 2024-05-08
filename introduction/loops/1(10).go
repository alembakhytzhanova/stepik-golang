package main

import "fmt"

// Функция для инвертирования порядка цифр в числе
func reverseNumber(num int) int {
	reversed := 0
	for num > 0 {
		digit := num % 10
		reversed = reversed*10 + digit
		num /= 10
	}
	return reversed
}

func main() {
	// Вводим два числа
	var num1, num2 int
	fmt.Scan(&num1, &num2)

	// Инвертируем порядок цифр в первом числе
	num1 = reverseNumber(num1)

	// Переменная для определения, является ли это первая цифра, чтобы не выводить пробел перед первой цифрой
	firstDigit := true

	// Перебираем цифры в первом числе
	for num1 > 0 {
		digit := num1 % 10
		num1 /= 10

		// Копируем второе число для каждой итерации цикла
		tempNum2 := num2

		// Проверяем второе число на наличие текущей цифры из первого числа
		for tempNum2 > 0 {
			if tempNum2%10 == digit {
				// Если цифра найдена и это не первая цифра, выводим пробел перед текущей цифрой
				if !firstDigit {
					fmt.Print(" ")
				}
				fmt.Print(digit)
				firstDigit = false
				break
			}
			tempNum2 /= 10
		}
	}

	// Если не найдено общих цифр, выводим "Нет общих цифр"
	if firstDigit {
		fmt.Println("Нет общих цифр")
	}
}
