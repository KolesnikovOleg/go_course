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
	for i := 1; i < len(A)+2; i++ {
		_, find := hashMap[i]
		if !find {
			return i
		}
	}

	fmt.Println("something happened wrong!!!")
	return -1
}

func main() {
	//numer fo array generation
	N := 10000
	A := make([]int, N)
	rand.Seed(time.Now().UnixNano())
	deleteId := rand.Intn(N) + 1
	for i := 1; i < N+1; i++ {
		if i != deleteId {
			A = append(A, i)
		}
	}
	fmt.Println("excepted value", deleteId)
	fmt.Println("Solution", Solution(A))
}
