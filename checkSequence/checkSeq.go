package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Solution(A []int) int {
	hashMap := make(map[int]int)
	for id, value := range A {
		hashMap[value] = id
	}
	for i := 0; i < len(A); i++ {
		_, find := hashMap[i+1]
		if !find {
			return 0
		}
	}

	return 1
}

func main() {
	//numer fo array generation
	N := 100
	A := []int{}
	rand.Seed(time.Now().UnixNano())
	needDelete := rand.Intn(2)
	deleteId := rand.Intn(N) + 1
	if needDelete > 0 {
		deleteId = -1
	}
	for i := 1; i < N+1; i++ {
		if i != deleteId {
			A = append(A, i)
		}
	}
	rand.Shuffle(len(A), func(i, j int) { A[i], A[j] = A[j], A[i] })
	fmt.Println("excepted value", needDelete, "LEN", len(A))
	fmt.Println("Solution", Solution(A))
}
