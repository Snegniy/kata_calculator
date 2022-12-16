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

func romanParse(rez int) {
	if rez < 1 {
		fmt.Println("Ошибка: римские цифры могут быть только положительными")
		os.Exit(1)
	}
	var output string
Loop:
	for rez > 0 {
		if rez == 100 {
			output += "C"
			rez -= 100
			continue Loop
		}
		if rez >= 90 {
			output += "XC"
			rez -= 90
			continue Loop
		}
		if rez >= 50 {
			output += "L"
			rez -= 50
			continue Loop
		}
		if rez >= 40 {
			output += "XL"
			rez -= 40
			continue Loop
		}
		if rez >= 10 {
			output += "X"
			rez -= 10
			continue Loop
		}
		if rez >= 9 {
			output += "IX"
			rez -= 9
			continue Loop
		}
		if rez >= 5 {
			output += "V"
			rez -= 5
			continue Loop
		}
		if rez >= 4 {
			output += "IV"
			rez -= 4
			continue Loop
		}
		if rez >= 1 {
			output += "I"
			rez -= 1
			continue Loop
		}
	}
	fmt.Println(output)
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
	if _, ok := romanNumbers[a]; ok {
		return romanNumbers[a]
	}
	return -1
}

func dataCheck(data []string) {
	var flag int
	if len(data) != 3 || (data[1] != "+" && data[1] != "-" && data[1] != "*" && data[1] != "/") {
		fmt.Println("Ошибка: неверный формат строки")
		os.Exit(1)
	}
	data[2] = strings.TrimSpace(data[2])
	a, err := strconv.Atoi(data[0])
	if err != nil {
		data[0] = strings.ToUpper(data[0])
		a = romanCheck(data[0])
		flag--
	} else {
		flag++
	}
	b, err := strconv.Atoi(data[2])
	if err != nil {
		data[2] = strings.ToUpper(data[2])
		b = romanCheck(data[2])
		flag--
	} else {
		flag++
	}
	if flag == 2 && (a >= 1 && a <= 10) && (b >= 1 && b <= 10) {
		fmt.Println(calculate(a, b, data[1]))
	} else if flag == -2 && (a >= 1 && a <= 10) && (b >= 1 && b <= 10) {
		romanParse(calculate(a, b, data[1]))
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
		data := strings.Split(text, " ")
		fmt.Println(len(data))
		dataCheck(data)
	}
}