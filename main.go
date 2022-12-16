package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculate(a, b int, c string) int {
	switch c {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		return 0
	}
}

func intToRoman(rez int) string {
	if rez < 1 {
		fmt.Println("Ошибка: римские цифры могут быть только положительными")
		os.Exit(1)
	}
	var output string

	for rez > 0 {
		if rez == 100 {
			output += "C"
			rez -= 100
		} else if rez >= 90 {
			output += "XC"
			rez -= 90
		} else if rez >= 50 {
			output += "L"
			rez -= 50
		} else if rez >= 40 {
			output += "XL"
			rez -= 40
		} else if rez >= 10 {
			output += "X"
			rez -= 10
		} else if rez >= 9 {
			output += "IX"
			rez -= 9
		} else if rez >= 5 {
			output += "V"
			rez -= 5
		} else if rez >= 4 {
			output += "IV"
			rez -= 4
		} else if rez >= 1 {
			output += "I"
			rez -= 1
		}
	}
	return output
}

func romanCheck(a string) int {
	var romanNumbers = map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}
	a = strings.ToUpper(a)
	if value, ok := romanNumbers[a]; ok {
		return value
	}
	return 0
}

func dataCheck(data []string) {
	var flag int
	if len(data) != 3 || (data[1] != "+" && data[1] != "-" && data[1] != "*" && data[1] != "/") {
		fmt.Println("Ошибка: неверный формат строки")
		os.Exit(1)
	}
	a, err := strconv.Atoi(data[0])
	if err != nil {
		a = romanCheck(data[0])
		flag--
	} else {
		flag++
	}
	b, err := strconv.Atoi(data[2])
	if err != nil {
		b = romanCheck(data[2])
		flag--
	} else {
		flag++
	}
	if flag == 2 && (a >= 1 && a <= 10) && (b >= 1 && b <= 10) {
		fmt.Println(calculate(a, b, data[1]))
	} else if flag == -2 && (a >= 1 && a <= 10) && (b >= 1 && b <= 10) {
		fmt.Println(intToRoman(calculate(a, b, data[1])))
	} else {
		fmt.Println("Ошибка: неверные данные")
		os.Exit(1)
	}
}

func main() {
	input := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Введите значения:")
		text, _ := input.ReadString('\n')
		data := strings.Fields(text)
		dataCheck(data)
	}
}
