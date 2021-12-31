package main

import (
	"fmt"
	"os"
)

func main() {
	var a int
	fmt.Println("Введите целое число:")
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("Число должно быть целым!")
		os.Exit(1)
	}
	for i := 2; i < a; i++{
		if a % i == 0 {
			fmt.Println("Число ", a, " не простое")
			os.Exit(1)
		}
	}
	fmt.Println("Число ", a, " простое")
}
