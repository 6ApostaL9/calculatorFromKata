package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Проверка на пустоту ввода
func checkEmptyInput(inputString string) {
	if len(inputString) == 0 {
		fmt.Println("Выдача паники, так как выражение не введено") //Обработка пустой строки
		os.Exit(1)
	}
}

// Проверка на корректность выражения
func checkCorrectExpression(arrInputString []string) {
	//Проверка длины выражения
	if len(arrInputString) != 3 {
		fmt.Println("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		os.Exit(2)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	inputString, _ := reader.ReadString('\n')    //Ждет ввода данных в формате строки
	inputString = strings.TrimSpace(inputString) //Очищает все пробелы

	checkEmptyInput(inputString)

	arrInputString := strings.Split(inputString, " ") //Преобразование строки в отдельные символы
	checkCorrectExpression(arrInputString)

	os.Exit(0)
}
