package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Solution(A []int) int {
	counter := len(A) + 1
	sum := 0
	for i := 0; i < len(A); i++ {
		counter += i + 1
		sum += A[i]
	}
	return counter - sum
}

func main() {
	//numer fo array generation
	N := 10000
	A := []int{}
	rand.Seed(time.Now().UnixNano())
	deleteId := rand.Intn(N) + 1
	for i := 1; i < N+1; i++ {
		if i != deleteId {
			A = append(A, i)
		}
	}
	rand.Shuffle(len(A), func(i, j int) { A[i], A[j] = A[j], A[i] })
	fmt.Println("excepted value", deleteId)
	fmt.Println("Solution", Solution(A))
}
