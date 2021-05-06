package main

import "fmt"

func hw3task1(n int) {
	//Задача1 Задача
	//Вывести на экран ряд чисел Фибоначчи, состоящий из n элементов.
	var ch, pCh, ppCh int

	for i := 0; i <= n; i++ {

		fmt.Printf("%v, число - %d \n", i, ch)

		if ch == 0 {
			ch = 1
			pCh = 0
			ppCh = 0
		} else {
			ppCh = pCh
			pCh = ch
			ch = pCh + ppCh
		}

	}
}

func hw3task2(n int) {
	//Задача 2 Напишите программу, доказывающую или проверяющую, что для множества натуральных чисел выполняется равенство:
	//1+2+...+n = n(n+1)/2, где n - любое натуральное число.
	var eq1 int
	for i := 1; i <= n; i++ {
		eq1 = eq1 + i
	}
	fmt.Printf("РАвенство %v = %v", eq1, n*(n+1)/2)

}
func hw3task3(n int) {
	var fak int = 1
	for i := 1; i <= n; i++ {
		fak = fak * i
	}
	fmt.Printf("Факториал числа %d! = %v\n", n, fak)
}
func main() {
	var numTask, inputTask1, inputTask2, inputTask3 int //
	fmt.Println("Введите номер задания.\n Задание 1 Вывод ряд Фибоначи. \n Задание 2. Доказать равенство 1+2+...+n = n(n+1)/2." +
		"\n Задание 3. Вычислить факториал введенного числа. ")
	fmt.Scan(&numTask)
	switch numTask {
	case 1:
		fmt.Println("Введите число ряда фибоначи\n")
		fmt.Scan(&inputTask1)
		hw3task1(inputTask1)
	case 2:
		fmt.Println("Введите число n \n")
		fmt.Scan(&inputTask2)
		hw3task2(inputTask2)
	case 3:
		println("Введите число n \n")
		fmt.Scan(&inputTask3)
		hw3task3(inputTask3)
	default:
		println("Please, enter  enter correct value")
	}

}
