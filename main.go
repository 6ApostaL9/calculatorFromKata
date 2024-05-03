package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const op = "+-*/"

func main() {
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	
	var x, a, y string
	//fmt.Scan(&x, &a, &y)
	fmt.Println(result(x, a, y))
	os.Exit(0)
}

// Результат выражения
func result(x, a, y string) string {
	parsSign(a)

	//Перевод строк ввода чисел в числа
	first, err1 := strconv.Atoi(x)
	second, err2 := strconv.Atoi(y)

	if err1 == nil && err2 == nil { //Если оба десятичные
		return strconv.Itoa(parsNumber(first, second, a))
	} else if (err1 != nil && err2 == nil) || (err1 == nil && err2 != nil) { //Если одно десятичное, а другое римское
		panic("Выдача паники, так как используются одновременно разные системы счисления.")
	} else { //Если оба римские
		return parsRome(x, a, y)
	}
}

// Проверка корректности ввода знака операции
func parsSign(a string) string {
	if strings.ContainsAny(a, op) {
		return a
	}
	panic("Выдача паники, так как строка не является математичской операцией.")
}

// Выполнение операций над числами
func parsNumber(x, y int, a string) int {
	var res int
	if x < 11 && y < 11 {
		switch a {
		case "+":
			res = x + y
		case "-":
			res = x - y
		case "*":
			res = x * y
		case "/":
			res = x / y
		}
		return res
	} else {
		panic("Выдача паники, так как формат математической опреации не соответствует заданию.")
	}

}

// Перевод римских чисел для выполнения операций
func parsRome(x, a, y string) string {
	numRome := map[string]int{ //Карта для перевода римских в десятичные
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
		"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}

	//Получение десятичных значений по ключу римских
	firstRome, ok1 := numRome[x]
	secondRome, ok2 := numRome[y]
	//Если удалось получить оба значения
	if ok1 && ok2 {
		num := parsNumber(firstRome, secondRome, a)
		if num < 0 {
			panic("Выдача паники, так как римской системе нет отрицательных чисел.")
		}

		//Перевод обратно в римскую систему
		//Выбраны такие числа, потому что в римской системе их написание отличается
		symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
		values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
		var result strings.Builder
		/*
			Вложенные циклы
			Находим какое число из десятичных меньше либо равно результату операции
			Добавляем римское написание к строке результата
			Повторяем пока число не закончится
		*/
		for i := 0; i < len(symbols); i++ {
			for num >= values[i] {
				result.WriteString(symbols[i])
				num -= values[i]
			}
		}
		return result.String()
	} else {
		panic("Выдача паники, так как формат математической операции не соответствует заданию.")
	}
}
