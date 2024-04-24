package main

import (
	"fmt"
	"wb_l2/dev_tasks/num2"
)

func main() {
	tests := []string{"ф4bc2ш5e", "abcd", "45", ""}
	for _, test := range tests {
		unpacked, err := num2.UnpackString(test)
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
		} else {
			fmt.Printf("Исходная строка: %s, Распакованная строка: %s\n", test, unpacked)
		}
	}

}
