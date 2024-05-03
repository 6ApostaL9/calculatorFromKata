package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	inputString, _ := reader.ReadString('\n')    //Ждет ввода данных в формате строки
	inputString = strings.TrimSpace(inputString) //Очищает все пробелы

	if len(inputString) == 0 {
		fmt.Println("Выдача паники, так как строка пустая") //Обработка пустой строки
		os.Exit(1)
	}

	os.Exit(0)
}
