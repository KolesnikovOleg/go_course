package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Solution(A []int) int {
	hashMap := make(map[int]int)
	for _, value := range A {
		hashMap[value]++

	}
	for key, value := range hashMap {
		if value%2 == 1 {
			return key
		}
	}

	fmt.Println("someone happened wrong!")
	return -1
}

func main() {
	//numer fo array generation
	N := 1000
	A := []int{}
	rand.Seed(time.Now().UnixNano())
	deleteId := rand.Intn(N)
	for i := 0; i < N; i += 2 {
		add := rand.Intn(N) + 1
		A = append(A, add, add)
	}
	availRes := A[deleteId]
	A = append(A[:deleteId], A[deleteId+1:]...)
	rand.Shuffle(len(A), func(i, j int) { A[i], A[j] = A[j], A[i] })
	fmt.Println("excepted value", availRes)
	fmt.Println("Solution", Solution(A))
}
