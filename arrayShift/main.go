package main

import "fmt"

func Solution(A []int, K int) []int {
	//Если K превышает длину массива, заменяем K на остаток от деления на длинну
	if K > len(A) {
		K = (K % len(A))
	}
	var startShift int = len(A) - K
	A = append(A[startShift:], A[:startShift]...)
	return A
}

func main() {
	// Заменить на нужжый массив для проверки
	A := []int{0, 1, 2, 3, 4, 5}
	// Заменить на нужное для проверки значение сдвига
	K := 8
	fmt.Println(Solution(A, K))
}
