package main

import "fmt"

/**
 * Для закрепления навыков отложенного вызова функций, напишите программу, содержащую вызов функции,
 * которая будет создавать паническую ситуацию неявно. Затем создайте отложенный вызов,
 * который будет обрабатывать эту паническую ситуацию и, в частности, печатать предупреждение в консоль.
 *
 * Критерием успешного выполнения задания является то, что программа не завершается аварийно ни при каких условиях.
 */

func main() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("Восстанавливаю работу после:", v)
		}
	}()
	funcWithPanic()
}

func funcWithPanic() {
	arr := []int{100, 300}

	for i := 0; i <= 2; i++ {
		fmt.Printf("Ключ: %d, значение: %d\n", i, arr[i])
	}
}
