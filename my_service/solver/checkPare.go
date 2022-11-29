package solver

import (
	"encoding/json"
	"fmt"
)

func NewCheckPareSolver(desc string) CheckPareSolver {
	return CheckPareSolver{
		description: desc,
	}
}

type CheckPareSolver struct {
	description string
}

func (c CheckPareSolver) Solution(A []int) int {
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

func (c CheckPareSolver) Solve(data interface{}) interface{} /*[][]int */ {
	res := []int{}
	for _, value := range data.([][][]int) {
		res = append(res, c.Solution(value[0]))
	}

	return res
}

func (c CheckPareSolver) Description() string {
	return c.description
}

func (c CheckPareSolver) Decode(rawData []byte) (interface{}, error) {
	var data [][][]int
	if err := json.Unmarshal(rawData, &data); err != nil {
		//fmt.Println(err)
		return nil, err
	}
	return data, nil
}
